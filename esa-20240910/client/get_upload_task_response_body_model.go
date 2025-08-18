// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUploadTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *GetUploadTaskResponseBody
	GetDescription() *string
	SetRequestId(v string) *GetUploadTaskResponseBody
	GetRequestId() *string
	SetStatus(v string) *GetUploadTaskResponseBody
	GetStatus() *string
}

type GetUploadTaskResponseBody struct {
	// The error message returned when the file upload task failed.
	//
	// example:
	//
	// invalid url
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ET5BF670-09D5-4D0B-BEBY-D96A2A52****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task status.
	//
	// 	- **INIT**: The task is being initialized.
	//
	// 	- **activacted**: The task is activated.
	//
	// 	- **running**: The task is running.
	//
	// 	- **success**: The task is successful.
	//
	// 	- **partial**: The task is partially successful.
	//
	// 	- **fail**: The task failed.
	//
	// example:
	//
	// running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetUploadTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUploadTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetUploadTaskResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetUploadTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUploadTaskResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetUploadTaskResponseBody) SetDescription(v string) *GetUploadTaskResponseBody {
	s.Description = &v
	return s
}

func (s *GetUploadTaskResponseBody) SetRequestId(v string) *GetUploadTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUploadTaskResponseBody) SetStatus(v string) *GetUploadTaskResponseBody {
	s.Status = &v
	return s
}

func (s *GetUploadTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
