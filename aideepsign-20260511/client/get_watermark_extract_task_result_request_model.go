// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWatermarkExtractTaskResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *GetWatermarkExtractTaskResultRequest
	GetClientToken() *string
	SetTaskId(v string) *GetWatermarkExtractTaskResultRequest
	GetTaskId() *string
}

type GetWatermarkExtractTaskResultRequest struct {
	// A reserved idempotence field. This field is not supported. Passing in a value does not cause an error, but it does not produce an idempotence effect.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF3898
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The dark watermark task ID returned by the CreateWatermarkExtractTask operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// f47ac10b-58cc-4372-a567-0e02b2c3d479
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetWatermarkExtractTaskResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWatermarkExtractTaskResultRequest) GoString() string {
	return s.String()
}

func (s *GetWatermarkExtractTaskResultRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *GetWatermarkExtractTaskResultRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetWatermarkExtractTaskResultRequest) SetClientToken(v string) *GetWatermarkExtractTaskResultRequest {
	s.ClientToken = &v
	return s
}

func (s *GetWatermarkExtractTaskResultRequest) SetTaskId(v string) *GetWatermarkExtractTaskResultRequest {
	s.TaskId = &v
	return s
}

func (s *GetWatermarkExtractTaskResultRequest) Validate() error {
	return dara.Validate(s)
}
