// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListInstancesResponseBodyData) *ListInstancesResponseBody
	GetData() []*ListInstancesResponseBodyData
	SetHttpCode(v int32) *ListInstancesResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *ListInstancesResponseBody
	GetRequestId() *string
}

type ListInstancesResponseBody struct {
	Data []*ListInstancesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// example:
	//
	// 0FC8BA40-C712-5FFD-9AA5-24C8F47F86E9
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBody) GetData() []*ListInstancesResponseBodyData {
	return s.Data
}

func (s *ListInstancesResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *ListInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstancesResponseBody) SetData(v []*ListInstancesResponseBodyData) *ListInstancesResponseBody {
	s.Data = v
	return s
}

func (s *ListInstancesResponseBody) SetHttpCode(v int32) *ListInstancesResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListInstancesResponseBody) SetRequestId(v string) *ListInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstancesResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInstancesResponseBodyData struct {
	// example:
	//
	// projectAbc
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListInstancesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListInstancesResponseBodyData) SetName(v string) *ListInstancesResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListInstancesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
