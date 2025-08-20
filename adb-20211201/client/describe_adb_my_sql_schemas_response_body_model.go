// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAdbMySqlSchemasResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *DescribeAdbMySqlSchemasResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeAdbMySqlSchemasResponseBody
	GetRequestId() *string
	SetSchemas(v []*string) *DescribeAdbMySqlSchemasResponseBody
	GetSchemas() []*string
	SetSuccess(v bool) *DescribeAdbMySqlSchemasResponseBody
	GetSuccess() *bool
}

type DescribeAdbMySqlSchemasResponseBody struct {
	// The returned message.
	//
	// 	- If the request was successful, a **success*	- message is returned.
	//
	// 	- If the request failed, an error message is returned.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The queried databases.
	Schemas []*string `json:"Schemas,omitempty" xml:"Schemas,omitempty" type:"Repeated"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAdbMySqlSchemasResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdbMySqlSchemasResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAdbMySqlSchemasResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeAdbMySqlSchemasResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAdbMySqlSchemasResponseBody) GetSchemas() []*string {
	return s.Schemas
}

func (s *DescribeAdbMySqlSchemasResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeAdbMySqlSchemasResponseBody) SetMessage(v string) *DescribeAdbMySqlSchemasResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAdbMySqlSchemasResponseBody) SetRequestId(v string) *DescribeAdbMySqlSchemasResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAdbMySqlSchemasResponseBody) SetSchemas(v []*string) *DescribeAdbMySqlSchemasResponseBody {
	s.Schemas = v
	return s
}

func (s *DescribeAdbMySqlSchemasResponseBody) SetSuccess(v bool) *DescribeAdbMySqlSchemasResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeAdbMySqlSchemasResponseBody) Validate() error {
	return dara.Validate(s)
}
