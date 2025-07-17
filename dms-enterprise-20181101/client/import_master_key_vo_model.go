// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportMasterKeyVO interface {
	dara.Model
	String() string
	GoString() string
	SetEncryptMekDataBase64(v string) *ImportMasterKeyVO
	GetEncryptMekDataBase64() *string
	SetMekId(v int64) *ImportMasterKeyVO
	GetMekId() *int64
	SetProjectId(v []*int64) *ImportMasterKeyVO
	GetProjectId() []*int64
}

type ImportMasterKeyVO struct {
	EncryptMekDataBase64 *string  `json:"EncryptMekDataBase64,omitempty" xml:"EncryptMekDataBase64,omitempty"`
	MekId                *int64   `json:"MekId,omitempty" xml:"MekId,omitempty"`
	ProjectId            []*int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty" type:"Repeated"`
}

func (s ImportMasterKeyVO) String() string {
	return dara.Prettify(s)
}

func (s ImportMasterKeyVO) GoString() string {
	return s.String()
}

func (s *ImportMasterKeyVO) GetEncryptMekDataBase64() *string {
	return s.EncryptMekDataBase64
}

func (s *ImportMasterKeyVO) GetMekId() *int64 {
	return s.MekId
}

func (s *ImportMasterKeyVO) GetProjectId() []*int64 {
	return s.ProjectId
}

func (s *ImportMasterKeyVO) SetEncryptMekDataBase64(v string) *ImportMasterKeyVO {
	s.EncryptMekDataBase64 = &v
	return s
}

func (s *ImportMasterKeyVO) SetMekId(v int64) *ImportMasterKeyVO {
	s.MekId = &v
	return s
}

func (s *ImportMasterKeyVO) SetProjectId(v []*int64) *ImportMasterKeyVO {
	s.ProjectId = v
	return s
}

func (s *ImportMasterKeyVO) Validate() error {
	return dara.Validate(s)
}
