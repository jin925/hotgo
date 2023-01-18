// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TestDao is the data access object for table hg_test.
type TestDao struct {
	table   string      // table is the underlying table name of the DAO.
	group   string      // group is the database configuration group name of current DAO.
	columns TestColumns // columns contains all the column names of Table for convenient usage.
}

// TestColumns defines and stores column names for table hg_test.
type TestColumns struct {
	Id          string // ID
	CategoryId  string // 分类ID
	Flag        string // 标签
	Title       string // 标题
	Description string // 描述
	Content     string // 内容
	Image       string // 单图
	Images      string // 多图
	Attachfile  string // 附件
	Attachfiles string // 多附件
	Map         string // 动态键值对
	Star        string // 推荐星
	Price       string // 价格
	Views       string // 浏览次数
	ActivityAt  string // 活动时间
	StartAt     string // 开启时间
	EndAt       string // 结束时间
	Switch      string // 开关
	Sort        string // 排序
	Avatar      string // 头像
	Sex         string // 性别
	Qq          string // qq
	Email       string // 邮箱
	Mobile      string // 手机号码
	Hobby       string // 爱好
	Channel     string // 渠道
	Pid         string // 上级ID
	Level       string // 树等级
	Tree        string // 关系树
	Remark      string // 备注
	Status      string // 状态
	CreatedBy   string // 创建者
	UpdatedBy   string // 更新者
	CreatedAt   string // 创建时间
	UpdatedAt   string // 修改时间
	DeletedAt   string // 删除时间
}

// testColumns holds the columns for table hg_test.
var testColumns = TestColumns{
	Id:          "id",
	CategoryId:  "category_id",
	Flag:        "flag",
	Title:       "title",
	Description: "description",
	Content:     "content",
	Image:       "image",
	Images:      "images",
	Attachfile:  "attachfile",
	Attachfiles: "attachfiles",
	Map:         "map",
	Star:        "star",
	Price:       "price",
	Views:       "views",
	ActivityAt:  "activity_at",
	StartAt:     "start_at",
	EndAt:       "end_at",
	Switch:      "switch",
	Sort:        "sort",
	Avatar:      "avatar",
	Sex:         "sex",
	Qq:          "qq",
	Email:       "email",
	Mobile:      "mobile",
	Hobby:       "hobby",
	Channel:     "channel",
	Pid:         "pid",
	Level:       "level",
	Tree:        "tree",
	Remark:      "remark",
	Status:      "status",
	CreatedBy:   "created_by",
	UpdatedBy:   "updated_by",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	DeletedAt:   "deleted_at",
}

// NewTestDao creates and returns a new DAO object for table data access.
func NewTestDao() *TestDao {
	return &TestDao{
		group:   "default",
		table:   "hg_test",
		columns: testColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TestDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TestDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TestDao) Columns() TestColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TestDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TestDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TestDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}