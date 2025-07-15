// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLiveEdgeTransferRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *SetLiveEdgeTransferRequest
	GetAppName() *string
	SetDomainName(v string) *SetLiveEdgeTransferRequest
	GetDomainName() *string
	SetHttpDns(v string) *SetLiveEdgeTransferRequest
	GetHttpDns() *string
	SetOwnerId(v int64) *SetLiveEdgeTransferRequest
	GetOwnerId() *int64
	SetRegionId(v string) *SetLiveEdgeTransferRequest
	GetRegionId() *string
	SetStreamName(v string) *SetLiveEdgeTransferRequest
	GetStreamName() *string
	SetTargetDomainList(v string) *SetLiveEdgeTransferRequest
	GetTargetDomainList() *string
	SetTransferArgs(v string) *SetLiveEdgeTransferRequest
	GetTransferArgs() *string
}

type SetLiveEdgeTransferRequest struct {
	// The name of the application to which the live stream belongs. Regular expressions are supported, with a few limits. For more information, see the **Description about the AppName and StreamName parameters*	- section. For example, a value of **liveApp\\*\\*\\*\\*[1,2,3]*	- specifies that stream relay is configured for three applications: liveApp\\*\\*\\*\\*1, liveApp\\*\\*\\*\\*2, and liveApp\\*\\*\\*\\*3.****
	//
	// >
	//
	// 	- This parameter takes effect for only destination domain names that are specified by the TargetDomainList parameter.
	//
	// 	- You cannot use a caret (^) or a dollar sign ($) in a regular expression to configure the `AppName` parameter. Otherwise, stream relay fails.
	//
	// example:
	//
	// liveApp****
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The ingest domain. Stream relay is configured based on the ingest domain. Only one stream relay configuration can be set for an ingest domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// demo.aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The HTTPDNS API that is used to obtain the destination URLs. The request must contain the `TargetDomainList` parameter or the `HttpDns` parameter. The two parameters are mutually exclusive.
	//
	// >  If the `HttpDns` parameter is configured, you cannot configure the `TargetDomainList` parameter, and the `AppName` and `StreamName` parameters do not take effect.
	//
	// For information about the requirements on the structure of messages that are returned by the HTTPDNS API, see the **Description about the HTTPDNS API*	- section.
	//
	// example:
	//
	// http://developer.aliyundoc.com
	HttpDns  *string `json:"HttpDns,omitempty" xml:"HttpDns,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the ingested stream. Regular expressions are supported, with a few limits. For more information, see the **Description about the AppName and StreamName parameters*	- section. For example, a value of **liveStream\\*\\*\\*\\*[1,2,3]*	- specifies that stream relay is configured for three streams: liveStream\\*\\*\\*\\*1, liveStream\\*\\*\\*\\*2, and liveStream\\*\\*\\*\\*3.****
	//
	// >
	//
	// 	- This parameter takes effect for only destination domain names that are specified by the TargetDomainList parameter.
	//
	// 	- You cannot use a caret (^) or a dollar sign ($) in a regular expression to configure the `StreamName` parameter. Otherwise, stream relay fails.
	//
	// example:
	//
	// liveStream****
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	// The destination domain names to which you want to relay the ingested stream. Separate multiple domain names with commas (,). The request must contain the `TargetDomainList` parameter or the `HttpDns` parameter. The two parameters are mutually exclusive.
	//
	// >
	//
	// 	- The `AppName` and `StreamName` parameters take effect only when the `TargetDomainList` parameter is configured.
	//
	// 	- If the `TargetDomainList` parameter is configured, you cannot configure the `HttpDns` parameter.
	//
	// example:
	//
	// learn.aliyundoc.com,guide.aliyundoc.com
	TargetDomainList *string `json:"TargetDomainList,omitempty" xml:"TargetDomainList,omitempty"`
	// Specifies whether to pass through stream ingest parameters. Valid values:
	//
	// 	- **yes**: passes through stream ingest parameters.
	//
	// 	- **no**: does not pass through stream ingest parameters.
	//
	// example:
	//
	// yes
	TransferArgs *string `json:"TransferArgs,omitempty" xml:"TransferArgs,omitempty"`
}

func (s SetLiveEdgeTransferRequest) String() string {
	return dara.Prettify(s)
}

func (s SetLiveEdgeTransferRequest) GoString() string {
	return s.String()
}

func (s *SetLiveEdgeTransferRequest) GetAppName() *string {
	return s.AppName
}

func (s *SetLiveEdgeTransferRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *SetLiveEdgeTransferRequest) GetHttpDns() *string {
	return s.HttpDns
}

func (s *SetLiveEdgeTransferRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SetLiveEdgeTransferRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SetLiveEdgeTransferRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *SetLiveEdgeTransferRequest) GetTargetDomainList() *string {
	return s.TargetDomainList
}

func (s *SetLiveEdgeTransferRequest) GetTransferArgs() *string {
	return s.TransferArgs
}

func (s *SetLiveEdgeTransferRequest) SetAppName(v string) *SetLiveEdgeTransferRequest {
	s.AppName = &v
	return s
}

func (s *SetLiveEdgeTransferRequest) SetDomainName(v string) *SetLiveEdgeTransferRequest {
	s.DomainName = &v
	return s
}

func (s *SetLiveEdgeTransferRequest) SetHttpDns(v string) *SetLiveEdgeTransferRequest {
	s.HttpDns = &v
	return s
}

func (s *SetLiveEdgeTransferRequest) SetOwnerId(v int64) *SetLiveEdgeTransferRequest {
	s.OwnerId = &v
	return s
}

func (s *SetLiveEdgeTransferRequest) SetRegionId(v string) *SetLiveEdgeTransferRequest {
	s.RegionId = &v
	return s
}

func (s *SetLiveEdgeTransferRequest) SetStreamName(v string) *SetLiveEdgeTransferRequest {
	s.StreamName = &v
	return s
}

func (s *SetLiveEdgeTransferRequest) SetTargetDomainList(v string) *SetLiveEdgeTransferRequest {
	s.TargetDomainList = &v
	return s
}

func (s *SetLiveEdgeTransferRequest) SetTransferArgs(v string) *SetLiveEdgeTransferRequest {
	s.TransferArgs = &v
	return s
}

func (s *SetLiveEdgeTransferRequest) Validate() error {
	return dara.Validate(s)
}
