// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeMeshVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpgradeMeshVersionResponseBody
	GetRequestId() *string
}

type UpgradeMeshVersionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 11fd0027-c27e-41bb-a565-75583054****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpgradeMeshVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpgradeMeshVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeMeshVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpgradeMeshVersionResponseBody) SetRequestId(v string) *UpgradeMeshVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeMeshVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
