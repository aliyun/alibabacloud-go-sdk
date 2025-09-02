// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTableModelInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateTableModelInfoResponseBody
	GetRequestId() *string
	SetUpdateResult(v bool) *UpdateTableModelInfoResponseBody
	GetUpdateResult() *bool
}

type UpdateTableModelInfoResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// abcde
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	UpdateResult *bool `json:"UpdateResult,omitempty" xml:"UpdateResult,omitempty"`
}

func (s UpdateTableModelInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTableModelInfoResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTableModelInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTableModelInfoResponseBody) GetUpdateResult() *bool {
	return s.UpdateResult
}

func (s *UpdateTableModelInfoResponseBody) SetRequestId(v string) *UpdateTableModelInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTableModelInfoResponseBody) SetUpdateResult(v bool) *UpdateTableModelInfoResponseBody {
	s.UpdateResult = &v
	return s
}

func (s *UpdateTableModelInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
