// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTranscodeTemplateGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNonExistTranscodeTemplateIds(v []*string) *DeleteTranscodeTemplateGroupResponseBody
	GetNonExistTranscodeTemplateIds() []*string
	SetRequestId(v string) *DeleteTranscodeTemplateGroupResponseBody
	GetRequestId() *string
}

type DeleteTranscodeTemplateGroupResponseBody struct {
	// The IDs of transcoding templates that were not found.
	NonExistTranscodeTemplateIds []*string `json:"NonExistTranscodeTemplateIds,omitempty" xml:"NonExistTranscodeTemplateIds,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 25818875-5F78-4A*****F6-D7393642CA58
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTranscodeTemplateGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTranscodeTemplateGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTranscodeTemplateGroupResponseBody) GetNonExistTranscodeTemplateIds() []*string {
	return s.NonExistTranscodeTemplateIds
}

func (s *DeleteTranscodeTemplateGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTranscodeTemplateGroupResponseBody) SetNonExistTranscodeTemplateIds(v []*string) *DeleteTranscodeTemplateGroupResponseBody {
	s.NonExistTranscodeTemplateIds = v
	return s
}

func (s *DeleteTranscodeTemplateGroupResponseBody) SetRequestId(v string) *DeleteTranscodeTemplateGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTranscodeTemplateGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
