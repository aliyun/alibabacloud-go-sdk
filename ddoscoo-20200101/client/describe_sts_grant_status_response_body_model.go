// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStsGrantStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeStsGrantStatusResponseBody
	GetRequestId() *string
	SetStsGrant(v *DescribeStsGrantStatusResponseBodyStsGrant) *DescribeStsGrantStatusResponseBody
	GetStsGrant() *DescribeStsGrantStatusResponseBodyStsGrant
}

type DescribeStsGrantStatusResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 6623EA1F-30FB-5BC8-BEC9-74D55F6F08F1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The authorization status of Anti-DDoS Pro or Anti-DDoS Premium.
	StsGrant *DescribeStsGrantStatusResponseBodyStsGrant `json:"StsGrant,omitempty" xml:"StsGrant,omitempty" type:"Struct"`
}

func (s DescribeStsGrantStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeStsGrantStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeStsGrantStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeStsGrantStatusResponseBody) GetStsGrant() *DescribeStsGrantStatusResponseBodyStsGrant {
	return s.StsGrant
}

func (s *DescribeStsGrantStatusResponseBody) SetRequestId(v string) *DescribeStsGrantStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeStsGrantStatusResponseBody) SetStsGrant(v *DescribeStsGrantStatusResponseBodyStsGrant) *DescribeStsGrantStatusResponseBody {
	s.StsGrant = v
	return s
}

func (s *DescribeStsGrantStatusResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeStsGrantStatusResponseBodyStsGrant struct {
	// The authorization status. Valid values:
	//
	// 	- **0**: Anti-DDoS Pro or Anti-DDoS Premium is not authorized to access other cloud services.
	//
	// 	- **1**: Anti-DDoS Pro or Anti-DDoS Premium is authorized to access other cloud services.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeStsGrantStatusResponseBodyStsGrant) String() string {
	return dara.Prettify(s)
}

func (s DescribeStsGrantStatusResponseBodyStsGrant) GoString() string {
	return s.String()
}

func (s *DescribeStsGrantStatusResponseBodyStsGrant) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeStsGrantStatusResponseBodyStsGrant) SetStatus(v int32) *DescribeStsGrantStatusResponseBodyStsGrant {
	s.Status = &v
	return s
}

func (s *DescribeStsGrantStatusResponseBodyStsGrant) Validate() error {
	return dara.Validate(s)
}
