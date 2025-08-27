// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainExceedApplyQueryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplyId(v int64) *TrainExceedApplyQueryRequest
	GetApplyId() *int64
	SetBusinessInstanceId(v string) *TrainExceedApplyQueryRequest
	GetBusinessInstanceId() *string
}

type TrainExceedApplyQueryRequest struct {
	// example:
	//
	// 349720
	ApplyId            *int64  `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	BusinessInstanceId *string `json:"business_instance_id,omitempty" xml:"business_instance_id,omitempty"`
}

func (s TrainExceedApplyQueryRequest) String() string {
	return dara.Prettify(s)
}

func (s TrainExceedApplyQueryRequest) GoString() string {
	return s.String()
}

func (s *TrainExceedApplyQueryRequest) GetApplyId() *int64 {
	return s.ApplyId
}

func (s *TrainExceedApplyQueryRequest) GetBusinessInstanceId() *string {
	return s.BusinessInstanceId
}

func (s *TrainExceedApplyQueryRequest) SetApplyId(v int64) *TrainExceedApplyQueryRequest {
	s.ApplyId = &v
	return s
}

func (s *TrainExceedApplyQueryRequest) SetBusinessInstanceId(v string) *TrainExceedApplyQueryRequest {
	s.BusinessInstanceId = &v
	return s
}

func (s *TrainExceedApplyQueryRequest) Validate() error {
	return dara.Validate(s)
}
