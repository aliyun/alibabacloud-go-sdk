// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMultimodeCmsUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *GetMultimodeCmsUrlResponseBody
	GetClusterId() *string
	SetMultimodCmsUrl(v string) *GetMultimodeCmsUrlResponseBody
	GetMultimodCmsUrl() *string
	SetRequestId(v string) *GetMultimodeCmsUrlResponseBody
	GetRequestId() *string
}

type GetMultimodeCmsUrlResponseBody struct {
	ClusterId      *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	MultimodCmsUrl *string `json:"MultimodCmsUrl,omitempty" xml:"MultimodCmsUrl,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMultimodeCmsUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMultimodeCmsUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetMultimodeCmsUrlResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetMultimodeCmsUrlResponseBody) GetMultimodCmsUrl() *string {
	return s.MultimodCmsUrl
}

func (s *GetMultimodeCmsUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMultimodeCmsUrlResponseBody) SetClusterId(v string) *GetMultimodeCmsUrlResponseBody {
	s.ClusterId = &v
	return s
}

func (s *GetMultimodeCmsUrlResponseBody) SetMultimodCmsUrl(v string) *GetMultimodeCmsUrlResponseBody {
	s.MultimodCmsUrl = &v
	return s
}

func (s *GetMultimodeCmsUrlResponseBody) SetRequestId(v string) *GetMultimodeCmsUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMultimodeCmsUrlResponseBody) Validate() error {
	return dara.Validate(s)
}
