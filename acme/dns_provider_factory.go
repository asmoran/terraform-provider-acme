// Auto-generated file. Do not edit.
package acme

import (
	"github.com/go-acme/lego/v4/challenge"
	"github.com/go-acme/lego/v4/providers/dns/acmedns"
	"github.com/go-acme/lego/v4/providers/dns/alidns"
	"github.com/go-acme/lego/v4/providers/dns/arvancloud"
	"github.com/go-acme/lego/v4/providers/dns/auroradns"
	"github.com/go-acme/lego/v4/providers/dns/autodns"
	"github.com/go-acme/lego/v4/providers/dns/azure"
	"github.com/go-acme/lego/v4/providers/dns/bindman"
	"github.com/go-acme/lego/v4/providers/dns/bluecat"
	"github.com/go-acme/lego/v4/providers/dns/checkdomain"
	"github.com/go-acme/lego/v4/providers/dns/clouddns"
	"github.com/go-acme/lego/v4/providers/dns/cloudflare"
	"github.com/go-acme/lego/v4/providers/dns/cloudns"
	"github.com/go-acme/lego/v4/providers/dns/cloudxns"
	"github.com/go-acme/lego/v4/providers/dns/conoha"
	"github.com/go-acme/lego/v4/providers/dns/constellix"
	"github.com/go-acme/lego/v4/providers/dns/desec"
	"github.com/go-acme/lego/v4/providers/dns/designate"
	"github.com/go-acme/lego/v4/providers/dns/digitalocean"
	"github.com/go-acme/lego/v4/providers/dns/dnsimple"
	"github.com/go-acme/lego/v4/providers/dns/dnsmadeeasy"
	"github.com/go-acme/lego/v4/providers/dns/dnspod"
	"github.com/go-acme/lego/v4/providers/dns/dode"
	"github.com/go-acme/lego/v4/providers/dns/domeneshop"
	"github.com/go-acme/lego/v4/providers/dns/dreamhost"
	"github.com/go-acme/lego/v4/providers/dns/duckdns"
	"github.com/go-acme/lego/v4/providers/dns/dyn"
	"github.com/go-acme/lego/v4/providers/dns/dynu"
	"github.com/go-acme/lego/v4/providers/dns/easydns"
	"github.com/go-acme/lego/v4/providers/dns/edgedns"
	"github.com/go-acme/lego/v4/providers/dns/exec"
	"github.com/go-acme/lego/v4/providers/dns/exoscale"
	"github.com/go-acme/lego/v4/providers/dns/gandi"
	"github.com/go-acme/lego/v4/providers/dns/gandiv5"
	"github.com/go-acme/lego/v4/providers/dns/gcloud"
	"github.com/go-acme/lego/v4/providers/dns/glesys"
	"github.com/go-acme/lego/v4/providers/dns/godaddy"
	"github.com/go-acme/lego/v4/providers/dns/hetzner"
	"github.com/go-acme/lego/v4/providers/dns/hostingde"
	"github.com/go-acme/lego/v4/providers/dns/httpreq"
	"github.com/go-acme/lego/v4/providers/dns/hurricane"
	"github.com/go-acme/lego/v4/providers/dns/hyperone"
	"github.com/go-acme/lego/v4/providers/dns/iij"
	"github.com/go-acme/lego/v4/providers/dns/infomaniak"
	"github.com/go-acme/lego/v4/providers/dns/inwx"
	"github.com/go-acme/lego/v4/providers/dns/ionos"
	"github.com/go-acme/lego/v4/providers/dns/joker"
	"github.com/go-acme/lego/v4/providers/dns/lightsail"
	"github.com/go-acme/lego/v4/providers/dns/linode"
	"github.com/go-acme/lego/v4/providers/dns/liquidweb"
	"github.com/go-acme/lego/v4/providers/dns/loopia"
	"github.com/go-acme/lego/v4/providers/dns/luadns"
	"github.com/go-acme/lego/v4/providers/dns/mydnsjp"
	"github.com/go-acme/lego/v4/providers/dns/mythicbeasts"
	"github.com/go-acme/lego/v4/providers/dns/namecheap"
	"github.com/go-acme/lego/v4/providers/dns/namedotcom"
	"github.com/go-acme/lego/v4/providers/dns/namesilo"
	"github.com/go-acme/lego/v4/providers/dns/netcup"
	"github.com/go-acme/lego/v4/providers/dns/netlify"
	"github.com/go-acme/lego/v4/providers/dns/nifcloud"
	"github.com/go-acme/lego/v4/providers/dns/njalla"
	"github.com/go-acme/lego/v4/providers/dns/ns1"
	"github.com/go-acme/lego/v4/providers/dns/oraclecloud"
	"github.com/go-acme/lego/v4/providers/dns/otc"
	"github.com/go-acme/lego/v4/providers/dns/ovh"
	"github.com/go-acme/lego/v4/providers/dns/pdns"
	"github.com/go-acme/lego/v4/providers/dns/rackspace"
	"github.com/go-acme/lego/v4/providers/dns/regru"
	"github.com/go-acme/lego/v4/providers/dns/rfc2136"
	"github.com/go-acme/lego/v4/providers/dns/rimuhosting"
	"github.com/go-acme/lego/v4/providers/dns/route53"
	"github.com/go-acme/lego/v4/providers/dns/sakuracloud"
	"github.com/go-acme/lego/v4/providers/dns/scaleway"
	"github.com/go-acme/lego/v4/providers/dns/selectel"
	"github.com/go-acme/lego/v4/providers/dns/servercow"
	"github.com/go-acme/lego/v4/providers/dns/stackpath"
	"github.com/go-acme/lego/v4/providers/dns/transip"
	"github.com/go-acme/lego/v4/providers/dns/vegadns"
	"github.com/go-acme/lego/v4/providers/dns/versio"
	"github.com/go-acme/lego/v4/providers/dns/vscale"
	"github.com/go-acme/lego/v4/providers/dns/vultr"
	"github.com/go-acme/lego/v4/providers/dns/yandex"
	"github.com/go-acme/lego/v4/providers/dns/zoneee"
	"github.com/go-acme/lego/v4/providers/dns/zonomi"
)

// dnsProviderFactoryFunc is a function that calls a provider's
// constructor and returns the provider interface.
type dnsProviderFactoryFunc func() (challenge.Provider, error)

// dnsProviderFactory is a factory for all of the valid DNS providers
// supported by ACME provider.
var dnsProviderFactory = map[string]dnsProviderFactoryFunc{
	"acme-dns": func() (challenge.Provider, error) {
		p, err := acmedns.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"alidns": func() (challenge.Provider, error) {
		p, err := alidns.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"arvancloud": func() (challenge.Provider, error) {
		p, err := arvancloud.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"auroradns": func() (challenge.Provider, error) {
		p, err := auroradns.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"autodns": func() (challenge.Provider, error) {
		p, err := autodns.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"azure": func() (challenge.Provider, error) {
		mapEnvironmentVariableValues(map[string]string{
			"ARM_CLIENT_ID":       "AZURE_CLIENT_ID",
			"ARM_CLIENT_SECRET":   "AZURE_CLIENT_SECRET",
			"ARM_RESOURCE_GROUP":  "AZURE_RESOURCE_GROUP",
			"ARM_SUBSCRIPTION_ID": "AZURE_SUBSCRIPTION_ID",
			"ARM_TENANT_ID":       "AZURE_TENANT_ID",
		})
		p, err := azure.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"bindman": func() (challenge.Provider, error) {
		p, err := bindman.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"bluecat": func() (challenge.Provider, error) {
		p, err := bluecat.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"checkdomain": func() (challenge.Provider, error) {
		p, err := checkdomain.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"clouddns": func() (challenge.Provider, error) {
		p, err := clouddns.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"cloudflare": func() (challenge.Provider, error) {
		p, err := cloudflare.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"cloudns": func() (challenge.Provider, error) {
		p, err := cloudns.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"cloudxns": func() (challenge.Provider, error) {
		p, err := cloudxns.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"conoha": func() (challenge.Provider, error) {
		p, err := conoha.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"constellix": func() (challenge.Provider, error) {
		p, err := constellix.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"desec": func() (challenge.Provider, error) {
		p, err := desec.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"designate": func() (challenge.Provider, error) {
		p, err := designate.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"digitalocean": func() (challenge.Provider, error) {
		p, err := digitalocean.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"dnsimple": func() (challenge.Provider, error) {
		p, err := dnsimple.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"dnsmadeeasy": func() (challenge.Provider, error) {
		p, err := dnsmadeeasy.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"dnspod": func() (challenge.Provider, error) {
		p, err := dnspod.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"dode": func() (challenge.Provider, error) {
		p, err := dode.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"domeneshop": func() (challenge.Provider, error) {
		p, err := domeneshop.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"dreamhost": func() (challenge.Provider, error) {
		p, err := dreamhost.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"duckdns": func() (challenge.Provider, error) {
		p, err := duckdns.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"dyn": func() (challenge.Provider, error) {
		p, err := dyn.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"dynu": func() (challenge.Provider, error) {
		p, err := dynu.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"easydns": func() (challenge.Provider, error) {
		p, err := easydns.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"edgedns": func() (challenge.Provider, error) {
		p, err := edgedns.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"exec": func() (challenge.Provider, error) {
		p, err := exec.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"exoscale": func() (challenge.Provider, error) {
		p, err := exoscale.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"gandi": func() (challenge.Provider, error) {
		p, err := gandi.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"gandiv5": func() (challenge.Provider, error) {
		p, err := gandiv5.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"gcloud": func() (challenge.Provider, error) {
		p, err := gcloud.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"glesys": func() (challenge.Provider, error) {
		p, err := glesys.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"godaddy": func() (challenge.Provider, error) {
		p, err := godaddy.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"hetzner": func() (challenge.Provider, error) {
		p, err := hetzner.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"hostingde": func() (challenge.Provider, error) {
		p, err := hostingde.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"httpreq": func() (challenge.Provider, error) {
		p, err := httpreq.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"hurricane": func() (challenge.Provider, error) {
		p, err := hurricane.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"hyperone": func() (challenge.Provider, error) {
		p, err := hyperone.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"iij": func() (challenge.Provider, error) {
		p, err := iij.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"infomaniak": func() (challenge.Provider, error) {
		p, err := infomaniak.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"inwx": func() (challenge.Provider, error) {
		p, err := inwx.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"ionos": func() (challenge.Provider, error) {
		p, err := ionos.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"joker": func() (challenge.Provider, error) {
		p, err := joker.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"lightsail": func() (challenge.Provider, error) {
		p, err := lightsail.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"linode": func() (challenge.Provider, error) {
		p, err := linode.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"liquidweb": func() (challenge.Provider, error) {
		p, err := liquidweb.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"loopia": func() (challenge.Provider, error) {
		p, err := loopia.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"luadns": func() (challenge.Provider, error) {
		p, err := luadns.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"mydnsjp": func() (challenge.Provider, error) {
		p, err := mydnsjp.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"mythicbeasts": func() (challenge.Provider, error) {
		p, err := mythicbeasts.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"namecheap": func() (challenge.Provider, error) {
		p, err := namecheap.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"namedotcom": func() (challenge.Provider, error) {
		p, err := namedotcom.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"namesilo": func() (challenge.Provider, error) {
		p, err := namesilo.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"netcup": func() (challenge.Provider, error) {
		p, err := netcup.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"netlify": func() (challenge.Provider, error) {
		p, err := netlify.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"nifcloud": func() (challenge.Provider, error) {
		p, err := nifcloud.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"njalla": func() (challenge.Provider, error) {
		p, err := njalla.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"ns1": func() (challenge.Provider, error) {
		p, err := ns1.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"oraclecloud": func() (challenge.Provider, error) {
		p, err := oraclecloud.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"otc": func() (challenge.Provider, error) {
		p, err := otc.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"ovh": func() (challenge.Provider, error) {
		p, err := ovh.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"pdns": func() (challenge.Provider, error) {
		p, err := pdns.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"rackspace": func() (challenge.Provider, error) {
		p, err := rackspace.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"regru": func() (challenge.Provider, error) {
		p, err := regru.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"rfc2136": func() (challenge.Provider, error) {
		p, err := rfc2136.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"rimuhosting": func() (challenge.Provider, error) {
		p, err := rimuhosting.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"route53": func() (challenge.Provider, error) {
		p, err := route53.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"sakuracloud": func() (challenge.Provider, error) {
		p, err := sakuracloud.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"scaleway": func() (challenge.Provider, error) {
		p, err := scaleway.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"selectel": func() (challenge.Provider, error) {
		p, err := selectel.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"servercow": func() (challenge.Provider, error) {
		p, err := servercow.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"stackpath": func() (challenge.Provider, error) {
		p, err := stackpath.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"transip": func() (challenge.Provider, error) {
		p, err := transip.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"vegadns": func() (challenge.Provider, error) {
		p, err := vegadns.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"versio": func() (challenge.Provider, error) {
		p, err := versio.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"vscale": func() (challenge.Provider, error) {
		p, err := vscale.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"vultr": func() (challenge.Provider, error) {
		p, err := vultr.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"yandex": func() (challenge.Provider, error) {
		p, err := yandex.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"zoneee": func() (challenge.Provider, error) {
		p, err := zoneee.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
	"zonomi": func() (challenge.Provider, error) {
		p, err := zonomi.NewDNSProvider()
		if err != nil {
			return nil, err
		}

		return p, nil
	},
}
