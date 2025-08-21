// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachSceneDefenseObjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AttachSceneDefenseObjectResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AttachSceneDefenseObjectResponseBody
	GetSuccess() *bool
}

type AttachSceneDefenseObjectResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// F65DF043-E0EB-4796-9467-23DDCDF92C1D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AttachSceneDefenseObjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AttachSceneDefenseObjectResponseBody) GoString() string {
	return s.String()
}

func (s *AttachSceneDefenseObjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AttachSceneDefenseObjectResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AttachSceneDefenseObjectResponseBody) SetRequestId(v string) *AttachSceneDefenseObjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachSceneDefenseObjectResponseBody) SetSuccess(v bool) *AttachSceneDefenseObjectResponseBody {
	s.Success = &v
	return s
}

func (s *AttachSceneDefenseObjectResponseBody) Validate() error {
	return dara.Validate(s)
}
