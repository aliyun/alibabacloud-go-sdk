// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMmsTablesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v int64) *UpdateMmsTablesResponseBody
	GetData() *int64
	SetRequestId(v string) *UpdateMmsTablesResponseBody
	GetRequestId() *string
}

type UpdateMmsTablesResponseBody struct {
	// example:
	//
	// 88
	Data *int64 `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// 5CA6292A-E301-5CD8-B4E2-AF060F99147B
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateMmsTablesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMmsTablesResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMmsTablesResponseBody) GetData() *int64 {
	return s.Data
}

func (s *UpdateMmsTablesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMmsTablesResponseBody) SetData(v int64) *UpdateMmsTablesResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateMmsTablesResponseBody) SetRequestId(v string) *UpdateMmsTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMmsTablesResponseBody) Validate() error {
	return dara.Validate(s)
}
