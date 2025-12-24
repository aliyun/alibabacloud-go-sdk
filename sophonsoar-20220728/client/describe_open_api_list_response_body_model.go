// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOpenApiListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeOpenApiListResponseBodyData) *DescribeOpenApiListResponseBody
	GetData() *DescribeOpenApiListResponseBodyData
	SetRequestId(v string) *DescribeOpenApiListResponseBody
	GetRequestId() *string
}

type DescribeOpenApiListResponseBody struct {
	// The data returned.
	Data *DescribeOpenApiListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// EF2ECA2D-D8E6-5021-BF5C-19DD6D52C5B2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeOpenApiListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeOpenApiListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOpenApiListResponseBody) GetData() *DescribeOpenApiListResponseBodyData {
	return s.Data
}

func (s *DescribeOpenApiListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeOpenApiListResponseBody) SetData(v *DescribeOpenApiListResponseBodyData) *DescribeOpenApiListResponseBody {
	s.Data = v
	return s
}

func (s *DescribeOpenApiListResponseBody) SetRequestId(v string) *DescribeOpenApiListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOpenApiListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeOpenApiListResponseBodyData struct {
	// The response code. Valid values:
	//
	// 	- 200: successful.
	//
	// 	- 500: failed.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the returned APIs.
	//
	// example:
	//
	// [{"apis":[{"summary":"get account information","deprecated":false,"name":"DescAccountSummary","title":"get account information"}],"childrens":[],"title":"account"}]
	Directories interface{} `json:"Directories,omitempty" xml:"Directories,omitempty"`
	// The version number of the API.
	//
	// example:
	//
	// 2018-12-03
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeOpenApiListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeOpenApiListResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeOpenApiListResponseBodyData) GetCode() *string {
	return s.Code
}

func (s *DescribeOpenApiListResponseBodyData) GetDirectories() interface{} {
	return s.Directories
}

func (s *DescribeOpenApiListResponseBodyData) GetVersion() *string {
	return s.Version
}

func (s *DescribeOpenApiListResponseBodyData) SetCode(v string) *DescribeOpenApiListResponseBodyData {
	s.Code = &v
	return s
}

func (s *DescribeOpenApiListResponseBodyData) SetDirectories(v interface{}) *DescribeOpenApiListResponseBodyData {
	s.Directories = v
	return s
}

func (s *DescribeOpenApiListResponseBodyData) SetVersion(v string) *DescribeOpenApiListResponseBodyData {
	s.Version = &v
	return s
}

func (s *DescribeOpenApiListResponseBodyData) Validate() error {
	return dara.Validate(s)
}
