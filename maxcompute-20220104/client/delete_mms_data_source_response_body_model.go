// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMmsDataSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v int64) *DeleteMmsDataSourceResponseBody
	GetData() *int64
	SetRequestId(v string) *DeleteMmsDataSourceResponseBody
	GetRequestId() *string
}

type DeleteMmsDataSourceResponseBody struct {
	// example:
	//
	// 2000015
	Data *int64 `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// EA1320AB-7766-5EC7-B0F6-8B20E2298567
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteMmsDataSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMmsDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMmsDataSourceResponseBody) GetData() *int64 {
	return s.Data
}

func (s *DeleteMmsDataSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMmsDataSourceResponseBody) SetData(v int64) *DeleteMmsDataSourceResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteMmsDataSourceResponseBody) SetRequestId(v string) *DeleteMmsDataSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMmsDataSourceResponseBody) Validate() error {
	return dara.Validate(s)
}
