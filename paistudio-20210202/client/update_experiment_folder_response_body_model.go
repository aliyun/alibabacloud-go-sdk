// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateExperimentFolderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateExperimentFolderResponseBody
	GetRequestId() *string
}

type UpdateExperimentFolderResponseBody struct {
	// example:
	//
	// E7C42CC7-2E85-508A-84F4-923B605FD10F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateExperimentFolderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateExperimentFolderResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateExperimentFolderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateExperimentFolderResponseBody) SetRequestId(v string) *UpdateExperimentFolderResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateExperimentFolderResponseBody) Validate() error {
	return dara.Validate(s)
}
