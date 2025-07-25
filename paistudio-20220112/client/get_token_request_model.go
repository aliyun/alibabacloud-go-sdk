// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExpireTime(v int64) *GetTokenRequest
	GetExpireTime() *int64
	SetTrainingJobId(v string) *GetTokenRequest
	GetTrainingJobId() *string
}

type GetTokenRequest struct {
	// example:
	//
	// 60
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// traincclrt205dcs
	TrainingJobId *string `json:"TrainingJobId,omitempty" xml:"TrainingJobId,omitempty"`
}

func (s GetTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTokenRequest) GoString() string {
	return s.String()
}

func (s *GetTokenRequest) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *GetTokenRequest) GetTrainingJobId() *string {
	return s.TrainingJobId
}

func (s *GetTokenRequest) SetExpireTime(v int64) *GetTokenRequest {
	s.ExpireTime = &v
	return s
}

func (s *GetTokenRequest) SetTrainingJobId(v string) *GetTokenRequest {
	s.TrainingJobId = &v
	return s
}

func (s *GetTokenRequest) Validate() error {
	return dara.Validate(s)
}
