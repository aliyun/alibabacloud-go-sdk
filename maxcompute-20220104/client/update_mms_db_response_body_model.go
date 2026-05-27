// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMmsDbResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v int64) *UpdateMmsDbResponseBody
	GetData() *int64
	SetRequestId(v string) *UpdateMmsDbResponseBody
	GetRequestId() *string
}

type UpdateMmsDbResponseBody struct {
	// example:
	//
	// success
	Data *int64 `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// 0b87b7e716665825896565060e87a4
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateMmsDbResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMmsDbResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMmsDbResponseBody) GetData() *int64 {
	return s.Data
}

func (s *UpdateMmsDbResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMmsDbResponseBody) SetData(v int64) *UpdateMmsDbResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateMmsDbResponseBody) SetRequestId(v string) *UpdateMmsDbResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMmsDbResponseBody) Validate() error {
	return dara.Validate(s)
}
