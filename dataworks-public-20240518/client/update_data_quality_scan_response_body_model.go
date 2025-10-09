// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataQualityScanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateDataQualityScanResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateDataQualityScanResponseBody
	GetSuccess() *bool
}

type UpdateDataQualityScanResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 22C97E95-F023-56B5-8852-B1A77A1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateDataQualityScanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataQualityScanResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDataQualityScanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDataQualityScanResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateDataQualityScanResponseBody) SetRequestId(v string) *UpdateDataQualityScanResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDataQualityScanResponseBody) SetSuccess(v bool) *UpdateDataQualityScanResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateDataQualityScanResponseBody) Validate() error {
	return dara.Validate(s)
}
