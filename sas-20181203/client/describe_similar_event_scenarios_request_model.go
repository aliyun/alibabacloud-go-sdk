// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSimilarEventScenariosRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceOwnerId(v int64) *DescribeSimilarEventScenariosRequest
	GetResourceOwnerId() *int64
	SetSecurityEventId(v int64) *DescribeSimilarEventScenariosRequest
	GetSecurityEventId() *int64
	SetSourceIp(v string) *DescribeSimilarEventScenariosRequest
	GetSourceIp() *string
}

type DescribeSimilarEventScenariosRequest struct {
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the alert event.
	//
	// >  You can call the [DescribeSuspEvents](~~DescribeSuspEvents~~) operation to query the ID of the alert event.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12321
	SecurityEventId *int64 `json:"SecurityEventId,omitempty" xml:"SecurityEventId,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 1.2.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeSimilarEventScenariosRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSimilarEventScenariosRequest) GoString() string {
	return s.String()
}

func (s *DescribeSimilarEventScenariosRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSimilarEventScenariosRequest) GetSecurityEventId() *int64 {
	return s.SecurityEventId
}

func (s *DescribeSimilarEventScenariosRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeSimilarEventScenariosRequest) SetResourceOwnerId(v int64) *DescribeSimilarEventScenariosRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSimilarEventScenariosRequest) SetSecurityEventId(v int64) *DescribeSimilarEventScenariosRequest {
	s.SecurityEventId = &v
	return s
}

func (s *DescribeSimilarEventScenariosRequest) SetSourceIp(v string) *DescribeSimilarEventScenariosRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeSimilarEventScenariosRequest) Validate() error {
	return dara.Validate(s)
}
