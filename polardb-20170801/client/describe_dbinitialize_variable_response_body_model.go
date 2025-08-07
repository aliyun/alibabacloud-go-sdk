// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInitializeVariableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBType(v string) *DescribeDBInitializeVariableResponseBody
	GetDBType() *string
	SetDBVersion(v string) *DescribeDBInitializeVariableResponseBody
	GetDBVersion() *string
	SetRequestId(v string) *DescribeDBInitializeVariableResponseBody
	GetRequestId() *string
	SetVariables(v *DescribeDBInitializeVariableResponseBodyVariables) *DescribeDBInitializeVariableResponseBody
	GetVariables() *DescribeDBInitializeVariableResponseBodyVariables
}

type DescribeDBInitializeVariableResponseBody struct {
	// The database type. Valid values:
	//
	// 	- Oracle
	//
	// 	- PostgreSQL
	//
	// 	- MySQL
	//
	// example:
	//
	// PostgreSQL
	DBType *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	// The version of the database engine.
	//
	// example:
	//
	// 11
	DBVersion *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 475F58B7-F394-4394-AA6E-4F1CBA******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The attributes that are returned.
	Variables *DescribeDBInitializeVariableResponseBodyVariables `json:"Variables,omitempty" xml:"Variables,omitempty" type:"Struct"`
}

func (s DescribeDBInitializeVariableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInitializeVariableResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInitializeVariableResponseBody) GetDBType() *string {
	return s.DBType
}

func (s *DescribeDBInitializeVariableResponseBody) GetDBVersion() *string {
	return s.DBVersion
}

func (s *DescribeDBInitializeVariableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBInitializeVariableResponseBody) GetVariables() *DescribeDBInitializeVariableResponseBodyVariables {
	return s.Variables
}

func (s *DescribeDBInitializeVariableResponseBody) SetDBType(v string) *DescribeDBInitializeVariableResponseBody {
	s.DBType = &v
	return s
}

func (s *DescribeDBInitializeVariableResponseBody) SetDBVersion(v string) *DescribeDBInitializeVariableResponseBody {
	s.DBVersion = &v
	return s
}

func (s *DescribeDBInitializeVariableResponseBody) SetRequestId(v string) *DescribeDBInitializeVariableResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInitializeVariableResponseBody) SetVariables(v *DescribeDBInitializeVariableResponseBodyVariables) *DescribeDBInitializeVariableResponseBody {
	s.Variables = v
	return s
}

func (s *DescribeDBInitializeVariableResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDBInitializeVariableResponseBodyVariables struct {
	Variable []*DescribeDBInitializeVariableResponseBodyVariablesVariable `json:"Variable,omitempty" xml:"Variable,omitempty" type:"Repeated"`
}

func (s DescribeDBInitializeVariableResponseBodyVariables) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInitializeVariableResponseBodyVariables) GoString() string {
	return s.String()
}

func (s *DescribeDBInitializeVariableResponseBodyVariables) GetVariable() []*DescribeDBInitializeVariableResponseBodyVariablesVariable {
	return s.Variable
}

func (s *DescribeDBInitializeVariableResponseBodyVariables) SetVariable(v []*DescribeDBInitializeVariableResponseBodyVariablesVariable) *DescribeDBInitializeVariableResponseBodyVariables {
	s.Variable = v
	return s
}

func (s *DescribeDBInitializeVariableResponseBodyVariables) Validate() error {
	return dara.Validate(s)
}

type DescribeDBInitializeVariableResponseBodyVariablesVariable struct {
	// The character set that is supported.
	//
	// example:
	//
	// EUC_CN
	Charset *string `json:"Charset,omitempty" xml:"Charset,omitempty"`
	// The language that indicates the collation of the databases that are created.
	//
	// >- The language must be compatible with the character set that is specified by **CharacterSetName**.
	//
	// >- This parameter is required for PolarDB for PostgreSQL (Compatible with Oracle) clusters or PolarDB for PostgreSQL clusters.
	//
	// >- This parameter is optional for PolarDB for MySQL clusters.
	//
	// To view the valid values for this parameter, perform the following steps: Log on to the PolarDB console and click the ID of a cluster. In the left-side navigation pane, choose **Settings and Management*	- > **Databases**. Then, click **Create Database**.
	//
	// example:
	//
	// C
	Collate *string `json:"Collate,omitempty" xml:"Collate,omitempty"`
	// The language that indicates the character type of the database.
	//
	// >
	//
	// 	- The language must be compatible with the character set that is specified by **CharacterSetName**.
	//
	// 	- The specified parameter value must be the same as the value of **Collate**.
	//
	// 	- If the PolarDB cluster runs PolarDB for PostgreSQL (Compatible with Oracle) or PolarDB for PostgreSQL, this parameter is required. If the cluster runs PolarDB for MySQL, this parameter is not supported.
	//
	// To view the valid values of this parameter, perform the following steps: First, log on to the PolarDB console and click the ID of a cluster. Then, in the left-side navigation pane, choose **Settings and Management*	- > **Databases**. Finally, click **Create Database**.
	//
	// example:
	//
	// C
	Ctype *string `json:"Ctype,omitempty" xml:"Ctype,omitempty"`
}

func (s DescribeDBInitializeVariableResponseBodyVariablesVariable) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInitializeVariableResponseBodyVariablesVariable) GoString() string {
	return s.String()
}

func (s *DescribeDBInitializeVariableResponseBodyVariablesVariable) GetCharset() *string {
	return s.Charset
}

func (s *DescribeDBInitializeVariableResponseBodyVariablesVariable) GetCollate() *string {
	return s.Collate
}

func (s *DescribeDBInitializeVariableResponseBodyVariablesVariable) GetCtype() *string {
	return s.Ctype
}

func (s *DescribeDBInitializeVariableResponseBodyVariablesVariable) SetCharset(v string) *DescribeDBInitializeVariableResponseBodyVariablesVariable {
	s.Charset = &v
	return s
}

func (s *DescribeDBInitializeVariableResponseBodyVariablesVariable) SetCollate(v string) *DescribeDBInitializeVariableResponseBodyVariablesVariable {
	s.Collate = &v
	return s
}

func (s *DescribeDBInitializeVariableResponseBodyVariablesVariable) SetCtype(v string) *DescribeDBInitializeVariableResponseBodyVariablesVariable {
	s.Ctype = &v
	return s
}

func (s *DescribeDBInitializeVariableResponseBodyVariablesVariable) Validate() error {
	return dara.Validate(s)
}
