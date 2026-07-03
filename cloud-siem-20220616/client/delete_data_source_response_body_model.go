// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DeleteDataSourceResponseBodyData) *DeleteDataSourceResponseBody
	GetData() *DeleteDataSourceResponseBodyData
	SetRequestId(v string) *DeleteDataSourceResponseBody
	GetRequestId() *string
}

type DeleteDataSourceResponseBody struct {
	// The return value of the request.
	Data *DeleteDataSourceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDataSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDataSourceResponseBody) GetData() *DeleteDataSourceResponseBodyData {
	return s.Data
}

func (s *DeleteDataSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDataSourceResponseBody) SetData(v *DeleteDataSourceResponseBodyData) *DeleteDataSourceResponseBody {
	s.Data = v
	return s
}

func (s *DeleteDataSourceResponseBody) SetRequestId(v string) *DeleteDataSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDataSourceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteDataSourceResponseBodyData struct {
	// The number of data sources that are deleted. A value of 1 indicates success. A value of 0 or less indicates failure.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DeleteDataSourceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataSourceResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteDataSourceResponseBodyData) GetCount() *int32 {
	return s.Count
}

func (s *DeleteDataSourceResponseBodyData) SetCount(v int32) *DeleteDataSourceResponseBodyData {
	s.Count = &v
	return s
}

func (s *DeleteDataSourceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
