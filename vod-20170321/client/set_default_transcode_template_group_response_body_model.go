// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDefaultTranscodeTemplateGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetDefaultTranscodeTemplateGroupResponseBody
	GetRequestId() *string
}

type SetDefaultTranscodeTemplateGroupResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 25818875-5F78-4A*****F6-D7393642CA58
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetDefaultTranscodeTemplateGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetDefaultTranscodeTemplateGroupResponseBody) GoString() string {
	return s.String()
}

func (s *SetDefaultTranscodeTemplateGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetDefaultTranscodeTemplateGroupResponseBody) SetRequestId(v string) *SetDefaultTranscodeTemplateGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetDefaultTranscodeTemplateGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
