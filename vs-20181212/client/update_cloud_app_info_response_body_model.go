// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloudAppInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPatchId(v string) *UpdateCloudAppInfoResponseBody
	GetPatchId() *string
	SetRequestId(v string) *UpdateCloudAppInfoResponseBody
	GetRequestId() *string
}

type UpdateCloudAppInfoResponseBody struct {
	// example:
	//
	// patch-03fa76e8e13a49b6a966b063d9d309b4
	PatchId *string `json:"PatchId,omitempty" xml:"PatchId,omitempty"`
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateCloudAppInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudAppInfoResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCloudAppInfoResponseBody) GetPatchId() *string {
	return s.PatchId
}

func (s *UpdateCloudAppInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCloudAppInfoResponseBody) SetPatchId(v string) *UpdateCloudAppInfoResponseBody {
	s.PatchId = &v
	return s
}

func (s *UpdateCloudAppInfoResponseBody) SetRequestId(v string) *UpdateCloudAppInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCloudAppInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
