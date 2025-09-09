// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseInstanceInternetAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *ReleaseInstanceInternetAddressResponseBody
	GetData() *bool
	SetRequestId(v string) *ReleaseInstanceInternetAddressResponseBody
	GetRequestId() *string
}

type ReleaseInstanceInternetAddressResponseBody struct {
	// The result returned by the current API.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// FD17CD3C-3355-49E8-9231-FE2DB0******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleaseInstanceInternetAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReleaseInstanceInternetAddressResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseInstanceInternetAddressResponseBody) GetData() *bool {
	return s.Data
}

func (s *ReleaseInstanceInternetAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReleaseInstanceInternetAddressResponseBody) SetData(v bool) *ReleaseInstanceInternetAddressResponseBody {
	s.Data = &v
	return s
}

func (s *ReleaseInstanceInternetAddressResponseBody) SetRequestId(v string) *ReleaseInstanceInternetAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReleaseInstanceInternetAddressResponseBody) Validate() error {
	return dara.Validate(s)
}
