// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCycleTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v string) *CreateCycleTaskResponseBody
	GetConfigId() *string
	SetRequestId(v string) *CreateCycleTaskResponseBody
	GetRequestId() *string
}

type CreateCycleTaskResponseBody struct {
	// The ID of the task configuration.
	//
	// example:
	//
	// 00cfa8161da093089e6804ba6a33****
	ConfigId *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 86CFF42E-E00A-57A3-8656-22291EFB****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCycleTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCycleTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCycleTaskResponseBody) GetConfigId() *string {
	return s.ConfigId
}

func (s *CreateCycleTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCycleTaskResponseBody) SetConfigId(v string) *CreateCycleTaskResponseBody {
	s.ConfigId = &v
	return s
}

func (s *CreateCycleTaskResponseBody) SetRequestId(v string) *CreateCycleTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCycleTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
