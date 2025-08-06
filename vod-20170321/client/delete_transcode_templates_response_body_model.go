// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTranscodeTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNonExistTranscodeTemplateIds(v []*string) *DeleteTranscodeTemplatesResponseBody
	GetNonExistTranscodeTemplateIds() []*string
	SetRequestId(v string) *DeleteTranscodeTemplatesResponseBody
	GetRequestId() *string
}

type DeleteTranscodeTemplatesResponseBody struct {
	NonExistTranscodeTemplateIds []*string `json:"NonExistTranscodeTemplateIds,omitempty" xml:"NonExistTranscodeTemplateIds,omitempty" type:"Repeated"`
	RequestId                    *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTranscodeTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTranscodeTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTranscodeTemplatesResponseBody) GetNonExistTranscodeTemplateIds() []*string {
	return s.NonExistTranscodeTemplateIds
}

func (s *DeleteTranscodeTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTranscodeTemplatesResponseBody) SetNonExistTranscodeTemplateIds(v []*string) *DeleteTranscodeTemplatesResponseBody {
	s.NonExistTranscodeTemplateIds = v
	return s
}

func (s *DeleteTranscodeTemplatesResponseBody) SetRequestId(v string) *DeleteTranscodeTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTranscodeTemplatesResponseBody) Validate() error {
	return dara.Validate(s)
}
