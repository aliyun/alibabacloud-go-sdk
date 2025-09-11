// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadClusterManagedCertResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *DownloadClusterManagedCertResponseBody
	GetContent() *string
	SetRequestId(v string) *DownloadClusterManagedCertResponseBody
	GetRequestId() *string
}

type DownloadClusterManagedCertResponseBody struct {
	// example:
	//
	// emhlbmdza****W5qaWFuYmlhbm1hY2VzaGk=
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 4C467B38-3910-447D-87BC-AC049166F216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DownloadClusterManagedCertResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DownloadClusterManagedCertResponseBody) GoString() string {
	return s.String()
}

func (s *DownloadClusterManagedCertResponseBody) GetContent() *string {
	return s.Content
}

func (s *DownloadClusterManagedCertResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DownloadClusterManagedCertResponseBody) SetContent(v string) *DownloadClusterManagedCertResponseBody {
	s.Content = &v
	return s
}

func (s *DownloadClusterManagedCertResponseBody) SetRequestId(v string) *DownloadClusterManagedCertResponseBody {
	s.RequestId = &v
	return s
}

func (s *DownloadClusterManagedCertResponseBody) Validate() error {
	return dara.Validate(s)
}
