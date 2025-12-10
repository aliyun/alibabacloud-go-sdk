// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteExperimentFolderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteExperimentFolderResponseBody
	GetRequestId() *string
}

type DeleteExperimentFolderResponseBody struct {
	// example:
	//
	// F2D0392B-D749-5C48-A98A-3FAE5C9444A6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteExperimentFolderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteExperimentFolderResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteExperimentFolderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteExperimentFolderResponseBody) SetRequestId(v string) *DeleteExperimentFolderResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteExperimentFolderResponseBody) Validate() error {
	return dara.Validate(s)
}
