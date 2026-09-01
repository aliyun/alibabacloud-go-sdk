// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSimilarSecurityEventsQueryTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateSimilarSecurityEventsQueryTaskRequest
	GetClientToken() *string
	SetResourceDirectoryAccountId(v int64) *CreateSimilarSecurityEventsQueryTaskRequest
	GetResourceDirectoryAccountId() *int64
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
	// The client token that is used to ensure the idempotence of the request. Different requests should use different tokens. The token supports only ASCII characters and cannot exceed 64 characters in length.
	ClientToken                *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ResourceDirectoryAccountId *int64  `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
	ResourceOwnerId            *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the security alert.
	//
	// This parameter is required.
	//
	// example:
	//
	// 14323
	SecurityEventId *int64 `json:"SecurityEventId,omitempty" xml:"SecurityEventId,omitempty"`
	// The code of the alerting event that has the same type or rule hits.
	//
	// example:
	//
	// default
	SimilarEventScenarioCode *string `json:"SimilarEventScenarioCode,omitempty" xml:"SimilarEventScenarioCode,omitempty"`
	// The IP address of the access source.
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

func (s *CreateSimilarSecurityEventsQueryTaskRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateSimilarSecurityEventsQueryTaskRequest) GetResourceDirectoryAccountId() *int64 {
	return s.ResourceDirectoryAccountId
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

func (s *CreateSimilarSecurityEventsQueryTaskRequest) SetClientToken(v string) *CreateSimilarSecurityEventsQueryTaskRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateSimilarSecurityEventsQueryTaskRequest) SetResourceDirectoryAccountId(v int64) *CreateSimilarSecurityEventsQueryTaskRequest {
	s.ResourceDirectoryAccountId = &v
	return s
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
