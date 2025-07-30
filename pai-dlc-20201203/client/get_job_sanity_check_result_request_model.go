// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobSanityCheckResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSanityCheckNumber(v int32) *GetJobSanityCheckResultRequest
	GetSanityCheckNumber() *int32
	SetSanityCheckPhase(v string) *GetJobSanityCheckResultRequest
	GetSanityCheckPhase() *string
	SetToken(v string) *GetJobSanityCheckResultRequest
	GetToken() *string
}

type GetJobSanityCheckResultRequest struct {
	// The nth time for which the job sanity check is performed.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	SanityCheckNumber *int32 `json:"SanityCheckNumber,omitempty" xml:"SanityCheckNumber,omitempty"`
	// The phase in which the job sanity check is performed.
	//
	// 	- CheckInit
	//
	// 	- DeviceCheck
	//
	// 	- SingleNodeCommCheck
	//
	// 	- TwoNodeCommCheck
	//
	// 	- AllNodeCommCheck
	//
	// example:
	//
	// DeviceCheck
	SanityCheckPhase *string `json:"SanityCheckPhase,omitempty" xml:"SanityCheckPhase,omitempty"`
	// The token information for job sharing. For more information about how to obtain the token information, see [GetToken](https://help.aliyun.com/document_detail/2557812.html).
	//
	// example:
	//
	// eyJhbG******zI1NiIsInR5cCI6IkpXVCJ9.eyJle****jE3MDk1Mzk0NDIsImlhdCI6MTcwODkzNDY0MiwidXNlcl9pZCI6IjE3NTgwNTQxNjI0Mzg2NTUiLCJ0YXJnZXRfaWQiOiJkbGM1OGh1a2xyYzZwdGMyIiwidGFyZ2V0X3R5cGUiOiJqb2IifQ.GNL7jo6****mgKKv0QeGIYgvBufSU-PH_EQttX****
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s GetJobSanityCheckResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetJobSanityCheckResultRequest) GoString() string {
	return s.String()
}

func (s *GetJobSanityCheckResultRequest) GetSanityCheckNumber() *int32 {
	return s.SanityCheckNumber
}

func (s *GetJobSanityCheckResultRequest) GetSanityCheckPhase() *string {
	return s.SanityCheckPhase
}

func (s *GetJobSanityCheckResultRequest) GetToken() *string {
	return s.Token
}

func (s *GetJobSanityCheckResultRequest) SetSanityCheckNumber(v int32) *GetJobSanityCheckResultRequest {
	s.SanityCheckNumber = &v
	return s
}

func (s *GetJobSanityCheckResultRequest) SetSanityCheckPhase(v string) *GetJobSanityCheckResultRequest {
	s.SanityCheckPhase = &v
	return s
}

func (s *GetJobSanityCheckResultRequest) SetToken(v string) *GetJobSanityCheckResultRequest {
	s.Token = &v
	return s
}

func (s *GetJobSanityCheckResultRequest) Validate() error {
	return dara.Validate(s)
}
