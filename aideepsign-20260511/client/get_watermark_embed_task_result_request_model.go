// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWatermarkEmbedTaskResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *GetWatermarkEmbedTaskResultRequest
	GetClientToken() *string
	SetTaskId(v string) *GetWatermarkEmbedTaskResultRequest
	GetTaskId() *string
}

type GetWatermarkEmbedTaskResultRequest struct {
	// A reserved idempotency field. This field is not supported. Passing a value does not cause an error, but it does not produce an idempotency effect.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF3898
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The invisible watermark task ID returned by the CreateWatermarkEmbedTask operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// f47ac10b-58cc-4372-a567-0e02b2c3d479
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetWatermarkEmbedTaskResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWatermarkEmbedTaskResultRequest) GoString() string {
	return s.String()
}

func (s *GetWatermarkEmbedTaskResultRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *GetWatermarkEmbedTaskResultRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetWatermarkEmbedTaskResultRequest) SetClientToken(v string) *GetWatermarkEmbedTaskResultRequest {
	s.ClientToken = &v
	return s
}

func (s *GetWatermarkEmbedTaskResultRequest) SetTaskId(v string) *GetWatermarkEmbedTaskResultRequest {
	s.TaskId = &v
	return s
}

func (s *GetWatermarkEmbedTaskResultRequest) Validate() error {
	return dara.Validate(s)
}
