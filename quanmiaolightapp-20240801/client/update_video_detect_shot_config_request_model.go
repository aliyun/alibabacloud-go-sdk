// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVideoDetectShotConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAsyncConcurrency(v int32) *UpdateVideoDetectShotConfigRequest
	GetAsyncConcurrency() *int32
}

type UpdateVideoDetectShotConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2
	AsyncConcurrency *int32 `json:"asyncConcurrency,omitempty" xml:"asyncConcurrency,omitempty"`
}

func (s UpdateVideoDetectShotConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateVideoDetectShotConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateVideoDetectShotConfigRequest) GetAsyncConcurrency() *int32 {
	return s.AsyncConcurrency
}

func (s *UpdateVideoDetectShotConfigRequest) SetAsyncConcurrency(v int32) *UpdateVideoDetectShotConfigRequest {
	s.AsyncConcurrency = &v
	return s
}

func (s *UpdateVideoDetectShotConfigRequest) Validate() error {
	return dara.Validate(s)
}
