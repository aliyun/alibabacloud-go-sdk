// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSubCrowdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateSubCrowdResponseBody
	GetRequestId() *string
	SetSubCrowdId(v string) *CreateSubCrowdResponseBody
	GetSubCrowdId() *string
}

type CreateSubCrowdResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 9763624B-5FBB-5E3A-9193-B1ADB554CEAE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 3
	SubCrowdId *string `json:"SubCrowdId,omitempty" xml:"SubCrowdId,omitempty"`
}

func (s CreateSubCrowdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSubCrowdResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSubCrowdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSubCrowdResponseBody) GetSubCrowdId() *string {
	return s.SubCrowdId
}

func (s *CreateSubCrowdResponseBody) SetRequestId(v string) *CreateSubCrowdResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSubCrowdResponseBody) SetSubCrowdId(v string) *CreateSubCrowdResponseBody {
	s.SubCrowdId = &v
	return s
}

func (s *CreateSubCrowdResponseBody) Validate() error {
	return dara.Validate(s)
}
