// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMmsJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v int64) *DeleteMmsJobResponseBody
	GetData() *int64
	SetRequestId(v string) *DeleteMmsJobResponseBody
	GetRequestId() *string
}

type DeleteMmsJobResponseBody struct {
	// example:
	//
	// 88
	Data *int64 `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// 31BE216F-AEF7-581E-B9C9-DECEB5424AC4
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteMmsJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMmsJobResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMmsJobResponseBody) GetData() *int64 {
	return s.Data
}

func (s *DeleteMmsJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMmsJobResponseBody) SetData(v int64) *DeleteMmsJobResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteMmsJobResponseBody) SetRequestId(v string) *DeleteMmsJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMmsJobResponseBody) Validate() error {
	return dara.Validate(s)
}
