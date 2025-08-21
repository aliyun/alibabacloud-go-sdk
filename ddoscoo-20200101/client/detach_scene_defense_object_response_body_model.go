// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachSceneDefenseObjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DetachSceneDefenseObjectResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DetachSceneDefenseObjectResponseBody
	GetSuccess() *bool
}

type DetachSceneDefenseObjectResponseBody struct {
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

func (s DetachSceneDefenseObjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetachSceneDefenseObjectResponseBody) GoString() string {
	return s.String()
}

func (s *DetachSceneDefenseObjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetachSceneDefenseObjectResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DetachSceneDefenseObjectResponseBody) SetRequestId(v string) *DetachSceneDefenseObjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachSceneDefenseObjectResponseBody) SetSuccess(v bool) *DetachSceneDefenseObjectResponseBody {
	s.Success = &v
	return s
}

func (s *DetachSceneDefenseObjectResponseBody) Validate() error {
	return dara.Validate(s)
}
