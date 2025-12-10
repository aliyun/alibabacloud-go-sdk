// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExperimentFolderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFolderId(v string) *CreateExperimentFolderResponseBody
	GetFolderId() *string
	SetRequestId(v string) *CreateExperimentFolderResponseBody
	GetRequestId() *string
}

type CreateExperimentFolderResponseBody struct {
	// example:
	//
	// folder-xxfdjhfxduxd
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// example:
	//
	// E7C42CC7-2E85-508A-84F4-923B605FD10F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateExperimentFolderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateExperimentFolderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateExperimentFolderResponseBody) GetFolderId() *string {
	return s.FolderId
}

func (s *CreateExperimentFolderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateExperimentFolderResponseBody) SetFolderId(v string) *CreateExperimentFolderResponseBody {
	s.FolderId = &v
	return s
}

func (s *CreateExperimentFolderResponseBody) SetRequestId(v string) *CreateExperimentFolderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateExperimentFolderResponseBody) Validate() error {
	return dara.Validate(s)
}
