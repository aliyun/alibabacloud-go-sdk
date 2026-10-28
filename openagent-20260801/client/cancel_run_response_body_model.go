// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelRunResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CancelRunResponseBodyData) *CancelRunResponseBody
	GetData() *CancelRunResponseBodyData
	SetRequestId(v string) *CancelRunResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CancelRunResponseBody
	GetSuccess() *bool
}

type CancelRunResponseBody struct {
	// The returned data.
	Data *CancelRunResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 00000000-0000-4000-8000-000000000001
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CancelRunResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelRunResponseBody) GoString() string {
	return s.String()
}

func (s *CancelRunResponseBody) GetData() *CancelRunResponseBodyData {
	return s.Data
}

func (s *CancelRunResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelRunResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CancelRunResponseBody) SetData(v *CancelRunResponseBodyData) *CancelRunResponseBody {
	s.Data = v
	return s
}

func (s *CancelRunResponseBody) SetRequestId(v string) *CancelRunResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelRunResponseBody) SetSuccess(v bool) *CancelRunResponseBody {
	s.Success = &v
	return s
}

func (s *CancelRunResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CancelRunResponseBodyData struct {
	// The run ID.
	//
	// example:
	//
	// run_example_001
	RunId *string `json:"runId,omitempty" xml:"runId,omitempty"`
	// The run status. If the cancellation is successful, the value is fixed to `CANCELLED`.
	//
	// example:
	//
	// CANCELLED
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s CancelRunResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CancelRunResponseBodyData) GoString() string {
	return s.String()
}

func (s *CancelRunResponseBodyData) GetRunId() *string {
	return s.RunId
}

func (s *CancelRunResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *CancelRunResponseBodyData) SetRunId(v string) *CancelRunResponseBodyData {
	s.RunId = &v
	return s
}

func (s *CancelRunResponseBodyData) SetStatus(v string) *CancelRunResponseBodyData {
	s.Status = &v
	return s
}

func (s *CancelRunResponseBodyData) Validate() error {
	return dara.Validate(s)
}
