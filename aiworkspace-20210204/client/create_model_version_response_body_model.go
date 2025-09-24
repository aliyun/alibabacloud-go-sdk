// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateModelVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateModelVersionResponseBody
	GetRequestId() *string
	SetVersionName(v string) *CreateModelVersionResponseBody
	GetVersionName() *string
}

type CreateModelVersionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 21645FCD-BAB9-5742-89AE-AEB27****B2E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The version of the model.
	//
	// example:
	//
	// 0.1.0
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s CreateModelVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateModelVersionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateModelVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateModelVersionResponseBody) GetVersionName() *string {
	return s.VersionName
}

func (s *CreateModelVersionResponseBody) SetRequestId(v string) *CreateModelVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateModelVersionResponseBody) SetVersionName(v string) *CreateModelVersionResponseBody {
	s.VersionName = &v
	return s
}

func (s *CreateModelVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
