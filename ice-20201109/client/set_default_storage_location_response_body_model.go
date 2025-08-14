// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDefaultStorageLocationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetDefaultStorageLocationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SetDefaultStorageLocationResponseBody
	GetSuccess() *bool
}

type SetDefaultStorageLocationResponseBody struct {
	// example:
	//
	// ******5A-CAAC-4850-A3AF-B74606******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetDefaultStorageLocationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetDefaultStorageLocationResponseBody) GoString() string {
	return s.String()
}

func (s *SetDefaultStorageLocationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetDefaultStorageLocationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SetDefaultStorageLocationResponseBody) SetRequestId(v string) *SetDefaultStorageLocationResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetDefaultStorageLocationResponseBody) SetSuccess(v bool) *SetDefaultStorageLocationResponseBody {
	s.Success = &v
	return s
}

func (s *SetDefaultStorageLocationResponseBody) Validate() error {
	return dara.Validate(s)
}
