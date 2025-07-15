// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveEdgeTransferResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *DescribeLiveEdgeTransferResponseBody
	GetAppName() *string
	SetDomainName(v string) *DescribeLiveEdgeTransferResponseBody
	GetDomainName() *string
	SetHttpDns(v string) *DescribeLiveEdgeTransferResponseBody
	GetHttpDns() *string
	SetRequestId(v string) *DescribeLiveEdgeTransferResponseBody
	GetRequestId() *string
	SetStreamName(v string) *DescribeLiveEdgeTransferResponseBody
	GetStreamName() *string
	SetTargetDomainList(v string) *DescribeLiveEdgeTransferResponseBody
	GetTargetDomainList() *string
	SetTransferArgs(v string) *DescribeLiveEdgeTransferResponseBody
	GetTransferArgs() *string
}

type DescribeLiveEdgeTransferResponseBody struct {
	// The name of the application to which the live stream belongs.
	//
	// example:
	//
	// liveApp****
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The ingest domain.
	//
	// example:
	//
	// demo.aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The HTTPDNS API that is used to obtain the destination URLs.
	//
	// example:
	//
	// http://developer.aliyundoc.com/****
	HttpDns *string `json:"HttpDns,omitempty" xml:"HttpDns,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 53FCB985-C67C-467B-B2B3-398430A21E14
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The name of the ingested stream.
	//
	// example:
	//
	// liveStream****
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	// The destination domain names to which the ingested stream is relayed. Multiple domain names are separated by commas (,).
	//
	// example:
	//
	// learn.aliyundoc.com,guide.aliyundoc.com
	TargetDomainList *string `json:"TargetDomainList,omitempty" xml:"TargetDomainList,omitempty"`
	// Indicates whether stream ingest parameters are passed through. Valid values:
	//
	// 	- **yes**: Stream ingest parameters are passed through.
	//
	// 	- **no**: Stream ingest parameters are not passed through.
	//
	// example:
	//
	// yes
	TransferArgs *string `json:"TransferArgs,omitempty" xml:"TransferArgs,omitempty"`
}

func (s DescribeLiveEdgeTransferResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveEdgeTransferResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveEdgeTransferResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *DescribeLiveEdgeTransferResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveEdgeTransferResponseBody) GetHttpDns() *string {
	return s.HttpDns
}

func (s *DescribeLiveEdgeTransferResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveEdgeTransferResponseBody) GetStreamName() *string {
	return s.StreamName
}

func (s *DescribeLiveEdgeTransferResponseBody) GetTargetDomainList() *string {
	return s.TargetDomainList
}

func (s *DescribeLiveEdgeTransferResponseBody) GetTransferArgs() *string {
	return s.TransferArgs
}

func (s *DescribeLiveEdgeTransferResponseBody) SetAppName(v string) *DescribeLiveEdgeTransferResponseBody {
	s.AppName = &v
	return s
}

func (s *DescribeLiveEdgeTransferResponseBody) SetDomainName(v string) *DescribeLiveEdgeTransferResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveEdgeTransferResponseBody) SetHttpDns(v string) *DescribeLiveEdgeTransferResponseBody {
	s.HttpDns = &v
	return s
}

func (s *DescribeLiveEdgeTransferResponseBody) SetRequestId(v string) *DescribeLiveEdgeTransferResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveEdgeTransferResponseBody) SetStreamName(v string) *DescribeLiveEdgeTransferResponseBody {
	s.StreamName = &v
	return s
}

func (s *DescribeLiveEdgeTransferResponseBody) SetTargetDomainList(v string) *DescribeLiveEdgeTransferResponseBody {
	s.TargetDomainList = &v
	return s
}

func (s *DescribeLiveEdgeTransferResponseBody) SetTransferArgs(v string) *DescribeLiveEdgeTransferResponseBody {
	s.TransferArgs = &v
	return s
}

func (s *DescribeLiveEdgeTransferResponseBody) Validate() error {
	return dara.Validate(s)
}
