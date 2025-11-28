// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIndexResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *CreateIndexResponseBody
	GetJobId() *string
	SetMessage(v string) *CreateIndexResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateIndexResponseBody
	GetRequestId() *string
	SetStatus(v string) *CreateIndexResponseBody
	GetStatus() *string
}

type CreateIndexResponseBody struct {
	// The job ID. It can be used to query the job status or cancel the job.
	//
	// example:
	//
	// 231460f8-75dc-405e-a669-0c5204887e91
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The returned message.
	//
	// example:
	//
	// Successfully create job
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique ID of the request.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the operation. Valid values:
	//
	// 	- **success**
	//
	// 	- **fail**
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateIndexResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateIndexResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIndexResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *CreateIndexResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateIndexResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateIndexResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateIndexResponseBody) SetJobId(v string) *CreateIndexResponseBody {
	s.JobId = &v
	return s
}

func (s *CreateIndexResponseBody) SetMessage(v string) *CreateIndexResponseBody {
	s.Message = &v
	return s
}

func (s *CreateIndexResponseBody) SetRequestId(v string) *CreateIndexResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateIndexResponseBody) SetStatus(v string) *CreateIndexResponseBody {
	s.Status = &v
	return s
}

func (s *CreateIndexResponseBody) Validate() error {
	return dara.Validate(s)
}
