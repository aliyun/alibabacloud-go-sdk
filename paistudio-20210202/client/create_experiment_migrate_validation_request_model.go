// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExperimentMigrateValidationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSourceExpId(v int64) *CreateExperimentMigrateValidationRequest
	GetSourceExpId() *int64
}

type CreateExperimentMigrateValidationRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123
	SourceExpId *int64 `json:"SourceExpId,omitempty" xml:"SourceExpId,omitempty"`
}

func (s CreateExperimentMigrateValidationRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateExperimentMigrateValidationRequest) GoString() string {
	return s.String()
}

func (s *CreateExperimentMigrateValidationRequest) GetSourceExpId() *int64 {
	return s.SourceExpId
}

func (s *CreateExperimentMigrateValidationRequest) SetSourceExpId(v int64) *CreateExperimentMigrateValidationRequest {
	s.SourceExpId = &v
	return s
}

func (s *CreateExperimentMigrateValidationRequest) Validate() error {
	return dara.Validate(s)
}
