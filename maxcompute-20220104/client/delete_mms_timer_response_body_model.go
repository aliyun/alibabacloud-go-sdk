// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMmsTimerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v int64) *DeleteMmsTimerResponseBody
	GetData() *int64
	SetRequestId(v string) *DeleteMmsTimerResponseBody
	GetRequestId() *string
}

type DeleteMmsTimerResponseBody struct {
	// example:
	//
	// success
	Data *int64 `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// 0a06dd4516687375802853481ec9fd
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteMmsTimerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMmsTimerResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMmsTimerResponseBody) GetData() *int64 {
	return s.Data
}

func (s *DeleteMmsTimerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMmsTimerResponseBody) SetData(v int64) *DeleteMmsTimerResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteMmsTimerResponseBody) SetRequestId(v string) *DeleteMmsTimerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMmsTimerResponseBody) Validate() error {
	return dara.Validate(s)
}
