// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkspaceAgenticFsMountRamAuthorizeUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessPointId(v string) *GetWorkspaceAgenticFsMountRamAuthorizeUrlRequest
	GetAccessPointId() *string
	SetFileSystemId(v string) *GetWorkspaceAgenticFsMountRamAuthorizeUrlRequest
	GetFileSystemId() *string
	SetServer(v string) *GetWorkspaceAgenticFsMountRamAuthorizeUrlRequest
	GetServer() *string
}

type GetWorkspaceAgenticFsMountRamAuthorizeUrlRequest struct {
	// The ID of the target NAS AccessPoint. This parameter corresponds to the server and fileSystemId parameters.
	//
	// This parameter is required.
	//
	// example:
	//
	// ap-0123456789abcdef0
	AccessPointId *string `json:"accessPointId,omitempty" xml:"accessPointId,omitempty"`
	// The ID of the NAS file system to which the target AccessPoint belongs. This parameter corresponds to the server and accessPointId parameters.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0123456789
	FileSystemId *string `json:"fileSystemId,omitempty" xml:"fileSystemId,omitempty"`
	// The domain name of the target AccessPoint, obtained from the DomainName field of the NAS ListAccessPoints operation. Do not include the protocol, port, or path.
	//
	// This parameter is required.
	//
	// example:
	//
	// ap-0123456789abcdef0.0123456789-vlm36.cn-hangzhou.nas.aliyuncs.com
	Server *string `json:"server,omitempty" xml:"server,omitempty"`
}

func (s GetWorkspaceAgenticFsMountRamAuthorizeUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspaceAgenticFsMountRamAuthorizeUrlRequest) GoString() string {
	return s.String()
}

func (s *GetWorkspaceAgenticFsMountRamAuthorizeUrlRequest) GetAccessPointId() *string {
	return s.AccessPointId
}

func (s *GetWorkspaceAgenticFsMountRamAuthorizeUrlRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *GetWorkspaceAgenticFsMountRamAuthorizeUrlRequest) GetServer() *string {
	return s.Server
}

func (s *GetWorkspaceAgenticFsMountRamAuthorizeUrlRequest) SetAccessPointId(v string) *GetWorkspaceAgenticFsMountRamAuthorizeUrlRequest {
	s.AccessPointId = &v
	return s
}

func (s *GetWorkspaceAgenticFsMountRamAuthorizeUrlRequest) SetFileSystemId(v string) *GetWorkspaceAgenticFsMountRamAuthorizeUrlRequest {
	s.FileSystemId = &v
	return s
}

func (s *GetWorkspaceAgenticFsMountRamAuthorizeUrlRequest) SetServer(v string) *GetWorkspaceAgenticFsMountRamAuthorizeUrlRequest {
	s.Server = &v
	return s
}

func (s *GetWorkspaceAgenticFsMountRamAuthorizeUrlRequest) Validate() error {
	return dara.Validate(s)
}
