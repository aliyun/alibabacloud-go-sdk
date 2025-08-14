// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExistSceneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeExistSceneResponseBody
	GetRequestId() *string
	SetData(v bool) *DescribeExistSceneResponseBody
	GetData() *bool
}

type DescribeExistSceneResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Data object
	//
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
}

func (s DescribeExistSceneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeExistSceneResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExistSceneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeExistSceneResponseBody) GetData() *bool {
	return s.Data
}

func (s *DescribeExistSceneResponseBody) SetRequestId(v string) *DescribeExistSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeExistSceneResponseBody) SetData(v bool) *DescribeExistSceneResponseBody {
	s.Data = &v
	return s
}

func (s *DescribeExistSceneResponseBody) Validate() error {
	return dara.Validate(s)
}
