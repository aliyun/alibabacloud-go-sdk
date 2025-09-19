// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSimilarSecurityEventsQueryTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceOwnerId(v int64) *CreateSimilarSecurityEventsQueryTaskRequest
	GetResourceOwnerId() *int64
	SetSecurityEventId(v int64) *CreateSimilarSecurityEventsQueryTaskRequest
	GetSecurityEventId() *int64
	SetSimilarEventScenarioCode(v string) *CreateSimilarSecurityEventsQueryTaskRequest
	GetSimilarEventScenarioCode() *string
	SetSourceIp(v string) *CreateSimilarSecurityEventsQueryTaskRequest
	GetSourceIp() *string
}

type CreateSimilarSecurityEventsQueryTaskRequest struct {
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the alert event.
	//
	// >  You must specify at least one of the SecurityEventId and SimilarEventScenarioCode parameters.
	//
	// This parameter is required.
	//
	// example:
	//
	// 14323
	SecurityEventId *int64 `json:"SecurityEventId,omitempty" xml:"SecurityEventId,omitempty"`
	// The codes of alert events that are triggered by the same rule or of the same alert type.
	//
	// >  You must specify at least one of the SecurityEventId and SimilarEventScenarioCode parameters.
	//
	// example:
	//
	// default
	SimilarEventScenarioCode *string `json:"SimilarEventScenarioCode,omitempty" xml:"SimilarEventScenarioCode,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 1.2.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s CreateSimilarSecurityEventsQueryTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSimilarSecurityEventsQueryTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateSimilarSecurityEventsQueryTaskRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateSimilarSecurityEventsQueryTaskRequest) GetSecurityEventId() *int64 {
	return s.SecurityEventId
}

func (s *CreateSimilarSecurityEventsQueryTaskRequest) GetSimilarEventScenarioCode() *string {
	return s.SimilarEventScenarioCode
}

func (s *CreateSimilarSecurityEventsQueryTaskRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *CreateSimilarSecurityEventsQueryTaskRequest) SetResourceOwnerId(v int64) *CreateSimilarSecurityEventsQueryTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateSimilarSecurityEventsQueryTaskRequest) SetSecurityEventId(v int64) *CreateSimilarSecurityEventsQueryTaskRequest {
	s.SecurityEventId = &v
	return s
}

func (s *CreateSimilarSecurityEventsQueryTaskRequest) SetSimilarEventScenarioCode(v string) *CreateSimilarSecurityEventsQueryTaskRequest {
	s.SimilarEventScenarioCode = &v
	return s
}

func (s *CreateSimilarSecurityEventsQueryTaskRequest) SetSourceIp(v string) *CreateSimilarSecurityEventsQueryTaskRequest {
	s.SourceIp = &v
	return s
}

func (s *CreateSimilarSecurityEventsQueryTaskRequest) Validate() error {
	return dara.Validate(s)
}
