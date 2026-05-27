// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMmsTableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v int64) *UpdateMmsTableResponseBody
	GetData() *int64
	SetRequestId(v string) *UpdateMmsTableResponseBody
	GetRequestId() *string
}

type UpdateMmsTableResponseBody struct {
	// example:
	//
	// 1
	Data *int64 `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// EA1320AB-7766-5EC7-B0F6-8B20E2298567
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateMmsTableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMmsTableResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMmsTableResponseBody) GetData() *int64 {
	return s.Data
}

func (s *UpdateMmsTableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMmsTableResponseBody) SetData(v int64) *UpdateMmsTableResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateMmsTableResponseBody) SetRequestId(v string) *UpdateMmsTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMmsTableResponseBody) Validate() error {
	return dara.Validate(s)
}
