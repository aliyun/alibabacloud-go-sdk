// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCrowdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCrowdId(v string) *CreateCrowdResponseBody
	GetCrowdId() *string
	SetRequestId(v string) *CreateCrowdResponseBody
	GetRequestId() *string
}

type CreateCrowdResponseBody struct {
	// example:
	//
	// 3
	CrowdId *string `json:"CrowdId,omitempty" xml:"CrowdId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 59CE7EC6-F268-5D71-9215-32922CC50D72
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCrowdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCrowdResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCrowdResponseBody) GetCrowdId() *string {
	return s.CrowdId
}

func (s *CreateCrowdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCrowdResponseBody) SetCrowdId(v string) *CreateCrowdResponseBody {
	s.CrowdId = &v
	return s
}

func (s *CreateCrowdResponseBody) SetRequestId(v string) *CreateCrowdResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCrowdResponseBody) Validate() error {
	return dara.Validate(s)
}
