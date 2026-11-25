// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitEditMediaManifestJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetManifestEditJobId(v string) *SubmitEditMediaManifestJobResponseBody
	GetManifestEditJobId() *string
	SetRequestId(v string) *SubmitEditMediaManifestJobResponseBody
	GetRequestId() *string
}

type SubmitEditMediaManifestJobResponseBody struct {
	// The manifest editing task ID, which is the MediaConvert pipeline task ID.
	//
	// example:
	//
	// 93e154cb83e54b7a0f1236f847f4****
	ManifestEditJobId *string `json:"ManifestEditJobId,omitempty" xml:"ManifestEditJobId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 25818875-5F78-4AF6-D7393642CA58****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitEditMediaManifestJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitEditMediaManifestJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitEditMediaManifestJobResponseBody) GetManifestEditJobId() *string {
	return s.ManifestEditJobId
}

func (s *SubmitEditMediaManifestJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitEditMediaManifestJobResponseBody) SetManifestEditJobId(v string) *SubmitEditMediaManifestJobResponseBody {
	s.ManifestEditJobId = &v
	return s
}

func (s *SubmitEditMediaManifestJobResponseBody) SetRequestId(v string) *SubmitEditMediaManifestJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitEditMediaManifestJobResponseBody) Validate() error {
	return dara.Validate(s)
}
