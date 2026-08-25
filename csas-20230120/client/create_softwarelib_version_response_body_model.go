// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSoftwarelibVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateSoftwarelibVersionResponseBody
	GetRequestId() *string
	SetVersionId(v string) *CreateSoftwarelibVersionResponseBody
	GetVersionId() *string
}

type CreateSoftwarelibVersionResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 2CABFEBB-0CE7-575E-833A-266F75D46713
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the software version that was created.
	//
	// example:
	//
	// softwarelib-version-21ae186e2ac9****
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s CreateSoftwarelibVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSoftwarelibVersionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSoftwarelibVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSoftwarelibVersionResponseBody) GetVersionId() *string {
	return s.VersionId
}

func (s *CreateSoftwarelibVersionResponseBody) SetRequestId(v string) *CreateSoftwarelibVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSoftwarelibVersionResponseBody) SetVersionId(v string) *CreateSoftwarelibVersionResponseBody {
	s.VersionId = &v
	return s
}

func (s *CreateSoftwarelibVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
