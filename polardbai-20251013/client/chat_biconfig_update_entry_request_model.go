// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatBIConfigUpdateEntryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbName(v string) *ChatBIConfigUpdateEntryRequest
	GetDbName() *string
	SetFormulaFunction(v string) *ChatBIConfigUpdateEntryRequest
	GetFormulaFunction() *string
	SetId(v int32) *ChatBIConfigUpdateEntryRequest
	GetId() *int32
	SetInstanceName(v string) *ChatBIConfigUpdateEntryRequest
	GetInstanceName() *string
	SetIsFunctional(v int32) *ChatBIConfigUpdateEntryRequest
	GetIsFunctional() *int32
	SetQueryFunction(v string) *ChatBIConfigUpdateEntryRequest
	GetQueryFunction() *string
	SetSqlCondition(v string) *ChatBIConfigUpdateEntryRequest
	GetSqlCondition() *string
	SetSqlFunction(v string) *ChatBIConfigUpdateEntryRequest
	GetSqlFunction() *string
	SetTextCondition(v string) *ChatBIConfigUpdateEntryRequest
	GetTextCondition() *string
}

type ChatBIConfigUpdateEntryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// db_test
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 若formula_function为总销售额：SUM(销售额)，则在最终处理时，问题中的总销售额将采用SUM(销售额)公式作为附加信息一并进行处理。
	FormulaFunction *string `json:"FormulaFunction,omitempty" xml:"FormulaFunction,omitempty"`
	// example:
	//
	// 1
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pc-2ze454l20me07****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// 1
	IsFunctional *int32 `json:"IsFunctional,omitempty" xml:"IsFunctional,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 若query_function为{"append":["一","二"],"delete":["？"],"replace":{"张三":"a","李四":"b"}}，则表示：当text_condition匹配时，在问题的结尾添加一和二，并删除问题中的？。最后，将问题中的张三替换为a，将李四替换为b。
	//
	// 例如：
	//
	// 问题张三今年总销售额是多少？：在text_condition匹配时，会最终被处理为a今年总销售额是多少一二。
	//
	// 问题李四今年总销售额多少？：在text_condition匹配时，会最终被处理为b今年总销售额是多少一二。
	QueryFunction *string `json:"QueryFunction,omitempty" xml:"QueryFunction,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 若sql_condition=students||student_courses&&!!courses，则表示：如果表students或表student_courses在SQL中，且表courses不在SQL中，则条件匹配。
	//
	// 例如：
	//
	// SQL语句SELECT 	- FROM student_courses：条件匹配。
	//
	// SQL语句SELECT c.course_name FROM student_courses sc JOIN courses c ON sc.courses_id = c.id;：条件不匹配。
	SqlCondition *string `json:"SqlCondition,omitempty" xml:"SqlCondition,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 若sql_function={"replace":{"status = \"请假\"":"status = 0","status = \"出勤\"":"status = 1"}}，则表示：在sql_condition匹配的情况下，将SQL中的status = \"请假\"替换为status = 0，status = \"出勤\"替换为status = 1。
	SqlFunction *string `json:"SqlFunction,omitempty" xml:"SqlFunction,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 若text_condition为张三||李四&&!!王五，则表示当问题包含张三，或者包含李四且不包含王五时，条件匹配。
	//
	// 例如：
	//
	// 张三今年总销售额多少？：条件匹配。
	//
	// 李四今年总销售额多少？：条件匹配。
	//
	// 李四王五今年总销售额多少？：条件不匹配。
	TextCondition *string `json:"TextCondition,omitempty" xml:"TextCondition,omitempty"`
}

func (s ChatBIConfigUpdateEntryRequest) String() string {
	return dara.Prettify(s)
}

func (s ChatBIConfigUpdateEntryRequest) GoString() string {
	return s.String()
}

func (s *ChatBIConfigUpdateEntryRequest) GetDbName() *string {
	return s.DbName
}

func (s *ChatBIConfigUpdateEntryRequest) GetFormulaFunction() *string {
	return s.FormulaFunction
}

func (s *ChatBIConfigUpdateEntryRequest) GetId() *int32 {
	return s.Id
}

func (s *ChatBIConfigUpdateEntryRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ChatBIConfigUpdateEntryRequest) GetIsFunctional() *int32 {
	return s.IsFunctional
}

func (s *ChatBIConfigUpdateEntryRequest) GetQueryFunction() *string {
	return s.QueryFunction
}

func (s *ChatBIConfigUpdateEntryRequest) GetSqlCondition() *string {
	return s.SqlCondition
}

func (s *ChatBIConfigUpdateEntryRequest) GetSqlFunction() *string {
	return s.SqlFunction
}

func (s *ChatBIConfigUpdateEntryRequest) GetTextCondition() *string {
	return s.TextCondition
}

func (s *ChatBIConfigUpdateEntryRequest) SetDbName(v string) *ChatBIConfigUpdateEntryRequest {
	s.DbName = &v
	return s
}

func (s *ChatBIConfigUpdateEntryRequest) SetFormulaFunction(v string) *ChatBIConfigUpdateEntryRequest {
	s.FormulaFunction = &v
	return s
}

func (s *ChatBIConfigUpdateEntryRequest) SetId(v int32) *ChatBIConfigUpdateEntryRequest {
	s.Id = &v
	return s
}

func (s *ChatBIConfigUpdateEntryRequest) SetInstanceName(v string) *ChatBIConfigUpdateEntryRequest {
	s.InstanceName = &v
	return s
}

func (s *ChatBIConfigUpdateEntryRequest) SetIsFunctional(v int32) *ChatBIConfigUpdateEntryRequest {
	s.IsFunctional = &v
	return s
}

func (s *ChatBIConfigUpdateEntryRequest) SetQueryFunction(v string) *ChatBIConfigUpdateEntryRequest {
	s.QueryFunction = &v
	return s
}

func (s *ChatBIConfigUpdateEntryRequest) SetSqlCondition(v string) *ChatBIConfigUpdateEntryRequest {
	s.SqlCondition = &v
	return s
}

func (s *ChatBIConfigUpdateEntryRequest) SetSqlFunction(v string) *ChatBIConfigUpdateEntryRequest {
	s.SqlFunction = &v
	return s
}

func (s *ChatBIConfigUpdateEntryRequest) SetTextCondition(v string) *ChatBIConfigUpdateEntryRequest {
	s.TextCondition = &v
	return s
}

func (s *ChatBIConfigUpdateEntryRequest) Validate() error {
	return dara.Validate(s)
}
