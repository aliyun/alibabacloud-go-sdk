// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSimilarEventScenariosRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSecurityEventId(v int64) *DescribeSimilarEventScenariosRequest
	GetSecurityEventId() *int64
}

type DescribeSimilarEventScenariosRequest struct {
	// The ID of the alert event.
	//
	// > You can call the [DescribeSuspEvents](~~DescribeSuspEvents~~) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12321
	SecurityEventId *int64 `json:"SecurityEventId,omitempty" xml:"SecurityEventId,omitempty"`
}

func (s DescribeSimilarEventScenariosRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSimilarEventScenariosRequest) GoString() string {
	return s.String()
}

func (s *DescribeSimilarEventScenariosRequest) GetSecurityEventId() *int64 {
	return s.SecurityEventId
}

func (s *DescribeSimilarEventScenariosRequest) SetSecurityEventId(v int64) *DescribeSimilarEventScenariosRequest {
	s.SecurityEventId = &v
	return s
}

func (s *DescribeSimilarEventScenariosRequest) Validate() error {
	return dara.Validate(s)
}
