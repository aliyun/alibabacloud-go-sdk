// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateQuotaPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *UpdateQuotaPlanResponseBody
	GetData() *string
	SetRequestId(v string) *UpdateQuotaPlanResponseBody
	GetRequestId() *string
}

type UpdateQuotaPlanResponseBody struct {
	// The returned result.
	//
	// example:
	//
	// success
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0a06dfe516688379832875789e2c65
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateQuotaPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateQuotaPlanResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateQuotaPlanResponseBody) GetData() *string {
	return s.Data
}

func (s *UpdateQuotaPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateQuotaPlanResponseBody) SetData(v string) *UpdateQuotaPlanResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateQuotaPlanResponseBody) SetRequestId(v string) *UpdateQuotaPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateQuotaPlanResponseBody) Validate() error {
	return dara.Validate(s)
}
