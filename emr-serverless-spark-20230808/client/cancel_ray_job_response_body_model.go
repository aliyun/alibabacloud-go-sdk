// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelRayJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CancelRayJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CancelRayJobResponseBody
	GetSuccess() *bool
}

type CancelRayJobResponseBody struct {
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CancelRayJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelRayJobResponseBody) GoString() string {
	return s.String()
}

func (s *CancelRayJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelRayJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CancelRayJobResponseBody) SetRequestId(v string) *CancelRayJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelRayJobResponseBody) SetSuccess(v bool) *CancelRayJobResponseBody {
	s.Success = &v
	return s
}

func (s *CancelRayJobResponseBody) Validate() error {
	return dara.Validate(s)
}
