// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMemoryAutoExtractionSwitchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClearedCount(v int64) *UpdateMemoryAutoExtractionSwitchResponseBody
	GetClearedCount() *int64
	SetCode(v string) *UpdateMemoryAutoExtractionSwitchResponseBody
	GetCode() *string
	SetEnabled(v bool) *UpdateMemoryAutoExtractionSwitchResponseBody
	GetEnabled() *bool
	SetMessage(v string) *UpdateMemoryAutoExtractionSwitchResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateMemoryAutoExtractionSwitchResponseBody
	GetRequestId() *string
}

type UpdateMemoryAutoExtractionSwitchResponseBody struct {
	// The number of automatic memory files cleared in this synchronous operation. The value is 0 if no files are cleared or no memories are available to be cleared.
	//
	// example:
	//
	// 3
	ClearedCount *int64 `json:"clearedCount,omitempty" xml:"clearedCount,omitempty"`
	// The business status code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The updated status of the automatic memory collection toggle. A value of true indicates that the feature is enabled, and a value of false indicates that the feature is disabled.
	//
	// example:
	//
	// false
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The error description.
	//
	// example:
	//
	// The current user is not an active member of the group cg_12345
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request tracking ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateMemoryAutoExtractionSwitchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMemoryAutoExtractionSwitchResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMemoryAutoExtractionSwitchResponseBody) GetClearedCount() *int64 {
	return s.ClearedCount
}

func (s *UpdateMemoryAutoExtractionSwitchResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateMemoryAutoExtractionSwitchResponseBody) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateMemoryAutoExtractionSwitchResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateMemoryAutoExtractionSwitchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMemoryAutoExtractionSwitchResponseBody) SetClearedCount(v int64) *UpdateMemoryAutoExtractionSwitchResponseBody {
	s.ClearedCount = &v
	return s
}

func (s *UpdateMemoryAutoExtractionSwitchResponseBody) SetCode(v string) *UpdateMemoryAutoExtractionSwitchResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateMemoryAutoExtractionSwitchResponseBody) SetEnabled(v bool) *UpdateMemoryAutoExtractionSwitchResponseBody {
	s.Enabled = &v
	return s
}

func (s *UpdateMemoryAutoExtractionSwitchResponseBody) SetMessage(v string) *UpdateMemoryAutoExtractionSwitchResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateMemoryAutoExtractionSwitchResponseBody) SetRequestId(v string) *UpdateMemoryAutoExtractionSwitchResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMemoryAutoExtractionSwitchResponseBody) Validate() error {
	return dara.Validate(s)
}
