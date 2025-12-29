// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddOrUpdateDisPlayModesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *AddOrUpdateDisPlayModesResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddOrUpdateDisPlayModesResponseBody
	GetRequestId() *string
	SetResult(v bool) *AddOrUpdateDisPlayModesResponseBody
	GetResult() *bool
	SetStatusCode(v int32) *AddOrUpdateDisPlayModesResponseBody
	GetStatusCode() *int32
}

type AddOrUpdateDisPlayModesResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 0EC7*726E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s AddOrUpdateDisPlayModesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddOrUpdateDisPlayModesResponseBody) GoString() string {
	return s.String()
}

func (s *AddOrUpdateDisPlayModesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddOrUpdateDisPlayModesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddOrUpdateDisPlayModesResponseBody) GetResult() *bool {
	return s.Result
}

func (s *AddOrUpdateDisPlayModesResponseBody) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddOrUpdateDisPlayModesResponseBody) SetMessage(v string) *AddOrUpdateDisPlayModesResponseBody {
	s.Message = &v
	return s
}

func (s *AddOrUpdateDisPlayModesResponseBody) SetRequestId(v string) *AddOrUpdateDisPlayModesResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddOrUpdateDisPlayModesResponseBody) SetResult(v bool) *AddOrUpdateDisPlayModesResponseBody {
	s.Result = &v
	return s
}

func (s *AddOrUpdateDisPlayModesResponseBody) SetStatusCode(v int32) *AddOrUpdateDisPlayModesResponseBody {
	s.StatusCode = &v
	return s
}

func (s *AddOrUpdateDisPlayModesResponseBody) Validate() error {
	return dara.Validate(s)
}
