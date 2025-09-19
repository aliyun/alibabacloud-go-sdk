// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOssScanConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *CreateOssScanConfigResponseBody
	GetId() *int64
	SetRequestId(v string) *CreateOssScanConfigResponseBody
	GetRequestId() *string
}

type CreateOssScanConfigResponseBody struct {
	// The policy ID.
	//
	// example:
	//
	// 210****
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 5DFD6277-CC36-57F7-ACE6-F5952123****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateOssScanConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateOssScanConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOssScanConfigResponseBody) GetId() *int64 {
	return s.Id
}

func (s *CreateOssScanConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateOssScanConfigResponseBody) SetId(v int64) *CreateOssScanConfigResponseBody {
	s.Id = &v
	return s
}

func (s *CreateOssScanConfigResponseBody) SetRequestId(v string) *CreateOssScanConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOssScanConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
