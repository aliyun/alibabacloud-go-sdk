// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFetchFieldsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateFetchFieldsResponseBody
	GetRequestId() *string
	SetResult(v bool) *UpdateFetchFieldsResponseBody
	GetResult() *bool
}

type UpdateFetchFieldsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D77D0DAF-790D-F5F5-A9C0-133738165014
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the operation was successful.
	//
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UpdateFetchFieldsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateFetchFieldsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFetchFieldsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateFetchFieldsResponseBody) GetResult() *bool {
	return s.Result
}

func (s *UpdateFetchFieldsResponseBody) SetRequestId(v string) *UpdateFetchFieldsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFetchFieldsResponseBody) SetResult(v bool) *UpdateFetchFieldsResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateFetchFieldsResponseBody) Validate() error {
	return dara.Validate(s)
}
