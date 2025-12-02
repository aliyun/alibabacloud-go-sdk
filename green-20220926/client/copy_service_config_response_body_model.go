// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyServiceConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *CopyServiceConfigResponseBody
	GetData() *bool
	SetRequestId(v string) *CopyServiceConfigResponseBody
	GetRequestId() *string
}

type CopyServiceConfigResponseBody struct {
	// The data returned.
	//
	// example:
	//
	// True
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CopyServiceConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CopyServiceConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CopyServiceConfigResponseBody) GetData() *bool {
	return s.Data
}

func (s *CopyServiceConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CopyServiceConfigResponseBody) SetData(v bool) *CopyServiceConfigResponseBody {
	s.Data = &v
	return s
}

func (s *CopyServiceConfigResponseBody) SetRequestId(v string) *CopyServiceConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *CopyServiceConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
