// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *GetJobStatusRequest
	GetClientToken() *string
	SetJobId(v string) *GetJobStatusRequest
	GetJobId() *string
}

type GetJobStatusRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// Generate a value from your client to ensure that the value is unique among different requests. ClientToken supports only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request is unique.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the asynchronous task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 72dcd26b-f12d-4c27-b3af-18f6aed5****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetJobStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetJobStatusRequest) GoString() string {
	return s.String()
}

func (s *GetJobStatusRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *GetJobStatusRequest) GetJobId() *string {
	return s.JobId
}

func (s *GetJobStatusRequest) SetClientToken(v string) *GetJobStatusRequest {
	s.ClientToken = &v
	return s
}

func (s *GetJobStatusRequest) SetJobId(v string) *GetJobStatusRequest {
	s.JobId = &v
	return s
}

func (s *GetJobStatusRequest) Validate() error {
	return dara.Validate(s)
}
