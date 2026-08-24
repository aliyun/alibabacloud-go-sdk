// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVirusFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDevTag(v string) *DeleteVirusFileRequest
	GetDevTag() *string
	SetFileMd5(v string) *DeleteVirusFileRequest
	GetFileMd5() *string
	SetFilePath(v string) *DeleteVirusFileRequest
	GetFilePath() *string
}

type DeleteVirusFileRequest struct {
	// The unique identifier of the user\\"s endpoint device where the virus file is located. The value can be up to 64 characters in length. You can obtain the value from the following operation:
	//
	// - [ListVirusFileStatuses](~~ListVirusFileStatuses~~): lists virus file statuses.
	//
	// This parameter is required.
	//
	// example:
	//
	// 36efa42d-2c32-c4dc-e3fc-8541e33a****
	DevTag *string `json:"DevTag,omitempty" xml:"DevTag,omitempty"`
	// The MD5 value of the virus file. The value must be a 32-character hexadecimal string. You can obtain the value from the following operation:
	//
	// - [ListVirusFileStatuses](~~ListVirusFileStatuses~~): lists virus file statuses.
	//
	// This parameter is required.
	//
	// example:
	//
	// d41d8cd98f00b204e9800998ecf8427e
	FileMd5 *string `json:"FileMd5,omitempty" xml:"FileMd5,omitempty"`
	// The full path of the virus file on the user\\"s endpoint device. Only records with a handling action of Fail can be deleted. You can obtain the value from the following operation:
	//
	// - [ListVirusFileStatuses](~~ListVirusFileStatuses~~): lists virus file statuses.
	//
	// This parameter is required.
	//
	// example:
	//
	// C:\\Users\\Public\\Downloads\\setup.exe
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
}

func (s DeleteVirusFileRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteVirusFileRequest) GoString() string {
	return s.String()
}

func (s *DeleteVirusFileRequest) GetDevTag() *string {
	return s.DevTag
}

func (s *DeleteVirusFileRequest) GetFileMd5() *string {
	return s.FileMd5
}

func (s *DeleteVirusFileRequest) GetFilePath() *string {
	return s.FilePath
}

func (s *DeleteVirusFileRequest) SetDevTag(v string) *DeleteVirusFileRequest {
	s.DevTag = &v
	return s
}

func (s *DeleteVirusFileRequest) SetFileMd5(v string) *DeleteVirusFileRequest {
	s.FileMd5 = &v
	return s
}

func (s *DeleteVirusFileRequest) SetFilePath(v string) *DeleteVirusFileRequest {
	s.FilePath = &v
	return s
}

func (s *DeleteVirusFileRequest) Validate() error {
	return dara.Validate(s)
}
