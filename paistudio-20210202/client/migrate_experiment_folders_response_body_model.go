// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMigrateExperimentFoldersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *MigrateExperimentFoldersResponseBody
	GetCode() *string
	SetFolderIdMapping(v map[string]map[string]interface{}) *MigrateExperimentFoldersResponseBody
	GetFolderIdMapping() map[string]map[string]interface{}
	SetMessage(v string) *MigrateExperimentFoldersResponseBody
	GetMessage() *string
	SetRequestId(v string) *MigrateExperimentFoldersResponseBody
	GetRequestId() *string
}

type MigrateExperimentFoldersResponseBody struct {
	// example:
	//
	// NO_PERMISSION
	Code            *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	FolderIdMapping map[string]map[string]interface{} `json:"FolderIdMapping,omitempty" xml:"FolderIdMapping,omitempty"`
	// example:
	//
	// SUCCESS
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9CFA2665-1FFE-5929-8468-C14C25890486
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MigrateExperimentFoldersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MigrateExperimentFoldersResponseBody) GoString() string {
	return s.String()
}

func (s *MigrateExperimentFoldersResponseBody) GetCode() *string {
	return s.Code
}

func (s *MigrateExperimentFoldersResponseBody) GetFolderIdMapping() map[string]map[string]interface{} {
	return s.FolderIdMapping
}

func (s *MigrateExperimentFoldersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *MigrateExperimentFoldersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MigrateExperimentFoldersResponseBody) SetCode(v string) *MigrateExperimentFoldersResponseBody {
	s.Code = &v
	return s
}

func (s *MigrateExperimentFoldersResponseBody) SetFolderIdMapping(v map[string]map[string]interface{}) *MigrateExperimentFoldersResponseBody {
	s.FolderIdMapping = v
	return s
}

func (s *MigrateExperimentFoldersResponseBody) SetMessage(v string) *MigrateExperimentFoldersResponseBody {
	s.Message = &v
	return s
}

func (s *MigrateExperimentFoldersResponseBody) SetRequestId(v string) *MigrateExperimentFoldersResponseBody {
	s.RequestId = &v
	return s
}

func (s *MigrateExperimentFoldersResponseBody) Validate() error {
	return dara.Validate(s)
}
