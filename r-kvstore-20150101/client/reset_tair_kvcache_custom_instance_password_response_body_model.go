// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetTairKVCacheCustomInstancePasswordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ResetTairKVCacheCustomInstancePasswordResponseBody
	GetRequestId() *string
}

type ResetTairKVCacheCustomInstancePasswordResponseBody struct {
	// example:
	//
	// AD425AD3-CC7B-4EE2-A5CB-2F61BA73****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetTairKVCacheCustomInstancePasswordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResetTairKVCacheCustomInstancePasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ResetTairKVCacheCustomInstancePasswordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResetTairKVCacheCustomInstancePasswordResponseBody) SetRequestId(v string) *ResetTairKVCacheCustomInstancePasswordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetTairKVCacheCustomInstancePasswordResponseBody) Validate() error {
	return dara.Validate(s)
}
