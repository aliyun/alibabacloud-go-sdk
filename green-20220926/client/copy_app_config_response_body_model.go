// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyAppConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *CopyAppConfigResponseBody
	GetData() *bool
	SetRequestId(v string) *CopyAppConfigResponseBody
	GetRequestId() *string
}

type CopyAppConfigResponseBody struct {
	// The returned data.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The ID assigned by the backend to uniquely identify the request. This ID can be used for troubleshooting.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CopyAppConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CopyAppConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CopyAppConfigResponseBody) GetData() *bool {
	return s.Data
}

func (s *CopyAppConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CopyAppConfigResponseBody) SetData(v bool) *CopyAppConfigResponseBody {
	s.Data = &v
	return s
}

func (s *CopyAppConfigResponseBody) SetRequestId(v string) *CopyAppConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *CopyAppConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
