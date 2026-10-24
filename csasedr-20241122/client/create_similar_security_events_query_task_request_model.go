// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSimilarSecurityEventsQueryTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSecurityEventId(v int64) *CreateSimilarSecurityEventsQueryTaskRequest
	GetSecurityEventId() *int64
	SetSimilarEventScenarioCode(v string) *CreateSimilarSecurityEventsQueryTaskRequest
	GetSimilarEventScenarioCode() *string
}

type CreateSimilarSecurityEventsQueryTaskRequest struct {
	// The ID of the alert event.
	//
	// This parameter is required.
	//
	// example:
	//
	// 14323
	SecurityEventId *int64 `json:"SecurityEventId,omitempty" xml:"SecurityEventId,omitempty"`
	// The scenario code of alerting events that are of the same type or have the same rule hits.
	//
	// example:
	//
	// default
	SimilarEventScenarioCode *string `json:"SimilarEventScenarioCode,omitempty" xml:"SimilarEventScenarioCode,omitempty"`
}

func (s CreateSimilarSecurityEventsQueryTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSimilarSecurityEventsQueryTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateSimilarSecurityEventsQueryTaskRequest) GetSecurityEventId() *int64 {
	return s.SecurityEventId
}

func (s *CreateSimilarSecurityEventsQueryTaskRequest) GetSimilarEventScenarioCode() *string {
	return s.SimilarEventScenarioCode
}

func (s *CreateSimilarSecurityEventsQueryTaskRequest) SetSecurityEventId(v int64) *CreateSimilarSecurityEventsQueryTaskRequest {
	s.SecurityEventId = &v
	return s
}

func (s *CreateSimilarSecurityEventsQueryTaskRequest) SetSimilarEventScenarioCode(v string) *CreateSimilarSecurityEventsQueryTaskRequest {
	s.SimilarEventScenarioCode = &v
	return s
}

func (s *CreateSimilarSecurityEventsQueryTaskRequest) Validate() error {
	return dara.Validate(s)
}
